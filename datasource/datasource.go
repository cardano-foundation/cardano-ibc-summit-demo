package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"

	"vesseloracle/x/vesseloracle/types"
)

const AddressPrefix = "vessel"                           // the address prefix of vessel oracle chain's addresses
const DataSourceCount = 8                                // number of emulated data sources
const DefaultVesselIMO = "9525338"                       // the default IMO of the vessel to fetch data for
const OutlierDeparturePortUnLoCode = "DEBWE"             // outlier departure port UNLOCODE identifier
const OutlierDeparturePortName = "BRAUNSCHWEIG"          // outlier departure port name
const EtaJitterOffsetSeconds int64 = 2 * 60              // eta jitter offset (2 minutes)
const EtaJitterIntervalWidthSeconds int64 = 4 * 60       // eta jitter interval (4 minutes)
const EtaOutlierOffsetSeconds int64 = 6 * 60 * 60        // minimum outlier difference for eta is 6 hours
const EtaOutlierIntervalWidthSeconds int64 = 4 * 60 * 60 // the outlier eta is in the interval of -[OutlierOffset, OutlierOffset+OutlierInterval)

// data type returned by datalastic vessel_pro endpoint
type DatalasticVesselPro struct {
	Data struct {
		UUID              string    `json:"uuid"`
		Name              string    `json:"name"`
		Mmsi              string    `json:"mmsi"`
		Imo               string    `json:"imo"`
		Eni               string    `json:"eni"`
		CountryIso        string    `json:"country_iso"`
		Type              string    `json:"type"`
		TypeSpecific      string    `json:"type_specific"`
		Lat               float64   `json:"lat"`
		Lon               float64   `json:"lon"`
		Speed             float64   `json:"speed"`
		Course            int32     `json:"course"`
		Heading           int32     `json:"heading"`
		CurrentDraught    int       `json:"current_draught"`
		NavigationStatus  string    `json:"navigation_status"`
		Destination       string    `json:"destination"`
		DestPortUUID      string    `json:"dest_port_uuid"`
		DestPort          string    `json:"dest_port"`
		DestPortUnlocode  string    `json:"dest_port_unlocode"`
		DepPortUUID       string    `json:"dep_port_uuid"`
		DepPort           string    `json:"dep_port"`
		DepPortUnlocode   string    `json:"dep_port_unlocode"`
		LastPositionEpoch int       `json:"last_position_epoch"`
		LastPositionUTC   time.Time `json:"last_position_UTC"`
		AtdEpoch          int64     `json:"atd_epoch"`
		AtdUTC            time.Time `json:"atd_UTC"`
		EtaEpoch          int64     `json:"eta_epoch"`
		EtaUTC            time.Time `json:"eta_UTC"`
		TimezoneOffsetSec int       `json:"timezone_offset_sec"`
		Timezone          string    `json:"timezone"`
	} `json:"data"`
	Meta struct {
		Duration float64 `json:"duration"`
		Endpoint string  `json:"endpoint"`
		Success  bool    `json:"success"`
	} `json:"meta"`
}

// fetch data from Datalastic API
func fetchVesselDataFromDatalastic(apiURL string, vesselImo string) (*DatalasticVesselPro, error) {
	resp, err := http.Get(apiURL) // Making an API call
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusForbidden {
			return nil, fmt.Errorf("No access to requested resource. Check the API keys. (%d)", resp.StatusCode)
		} else if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("Vessel with given IMO not found. (%s)", vesselImo)
		} else {
			return nil, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var vesselData DatalasticVesselPro
	json.Unmarshal(body, &vesselData)

	return &vesselData, nil
}

// create data sets for vessels including outliers for ETA and departure port
func generateVesselData(apiData *DatalasticVesselPro) ([]DatalasticVesselPro, error) {
	var vesselDataSources [DataSourceCount]DatalasticVesselPro
	for vesselDataSourceIndex := 0; vesselDataSourceIndex < DataSourceCount; vesselDataSourceIndex++ {
		vesselDataSources[vesselDataSourceIndex] = *apiData

		// only leave the first data source unchanged compared to the reference and introduce slight eta jitter for every other data source, except the outlier
		if vesselDataSourceIndex > 1 {
			vesselDataSources[vesselDataSourceIndex].Data.EtaEpoch = vesselDataSources[vesselDataSourceIndex].Data.EtaEpoch - EtaJitterOffsetSeconds + rand.Int63n(EtaJitterIntervalWidthSeconds)
			vesselDataSources[vesselDataSourceIndex].Data.EtaUTC = time.Unix(int64(vesselDataSources[vesselDataSourceIndex].Data.EtaEpoch), 0).UTC()
		}

		// create outliers: for index == 1 create the eta outlier, for index == 2 create the departure port outlier
		if vesselDataSourceIndex == 1 {
			vesselDataSources[vesselDataSourceIndex].Data.EtaEpoch = vesselDataSources[vesselDataSourceIndex].Data.EtaEpoch - EtaOutlierOffsetSeconds - rand.Int63n(EtaOutlierIntervalWidthSeconds)
			vesselDataSources[vesselDataSourceIndex].Data.EtaUTC = time.Unix(int64(vesselDataSources[vesselDataSourceIndex].Data.EtaEpoch), 0).UTC()
		} else if vesselDataSourceIndex == 2 {
			vesselDataSources[vesselDataSourceIndex].Data.DepPortUnlocode = OutlierDeparturePortUnLoCode
			vesselDataSources[vesselDataSourceIndex].Data.DepPort = OutlierDeparturePortName
		}
	}
	return vesselDataSources[:], nil
}

// create transactions based on vessel data and submit them, returning the list of tx identifiers for further processing
func createAndSubmitDataReports(vesselDataSources []DatalasticVesselPro) ([]string, error) {
	var transactions []string = make([]string, len(vesselDataSources))

	ctx := context.Background()

	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(AddressPrefix))
	if err != nil {
		return transactions, fmt.Errorf("Could not create cosmos client: %v", err)
	}

	for index, vesselData := range vesselDataSources {
		dataSourceAccountName := fmt.Sprintf("ds%d", index)
		dataSourceAccount, err := client.Account(dataSourceAccountName)
		if err != nil {
			return transactions, fmt.Errorf("Could not determine account for data source: %v %v", dataSourceAccountName, err)
		}

		dataSourceAddress, err := dataSourceAccount.Address(AddressPrefix)
		if err != nil {
			return transactions, fmt.Errorf("Could not determine address for account: %v %v", dataSourceAccountName, err)
		}

		// Define a message to create a post
		msg := &types.MsgCreateVessel{
			Creator:  dataSourceAddress,
			Imo:      vesselData.Data.Imo,
			Ts:       uint64(vesselData.Data.LastPositionEpoch),
			Source:   dataSourceAddress,
			Lat:      int32(vesselData.Data.Lat * 100000), // convert to fixed point representation, multiply by 100000
			Lon:      int32(vesselData.Data.Lon * 100000), // convert to fixed point representation, multiply by 100000
			Speed:    int32(vesselData.Data.Speed * 10),   // convert to fixed point representation, multiply by 10
			Course:   vesselData.Data.Course,
			Heading:  vesselData.Data.Heading,
			Adt:      uint64(vesselData.Data.AtdEpoch),
			Eta:      uint64(vesselData.Data.EtaEpoch),
			Name:     vesselData.Data.Name,
			Destport: vesselData.Data.DestPortUnlocode,
			Depport:  vesselData.Data.DepPortUnlocode,
			Mmsi:     vesselData.Data.Mmsi,
		}

		fmt.Println("Submitting data report for", dataSourceAccountName, "with address", dataSourceAddress, "...")
		txResp, err := client.BroadcastTx(ctx, dataSourceAccount, msg)
		if err != nil {
			return transactions, fmt.Errorf("Could not broadcast transaction for data source %v: %v", dataSourceAccountName, err)
		}
		transactions[index] = txResp.TxHash
		fmt.Println(txResp)
	}

	return transactions, nil
}

// submit a data consolidation request message
func submitDataConsolidationRequest(imo string) (*string, error) {
	ctx := context.Background()

	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(AddressPrefix))
	if err != nil {
		return nil, fmt.Errorf("Could not create cosmos client: %v", err)
	}

	account, err := client.Account("bob")
	if err != nil {
		return nil, fmt.Errorf("Could not determine account bob: %v", err)
	}

	address, err := account.Address(AddressPrefix)
	if err != nil {
		return nil, fmt.Errorf("Could not determine address for account bob: %v", err)
	}

	msg := &types.MsgConsolidateReports{
		Creator: address,
		Imo:     imo,
	}

	fmt.Println("Submitting request for data consolidation from account bob, with address", address, "...")
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		return nil, fmt.Errorf("Could not broadcast transaction for data consolidation request: %v", err)
	}
	fmt.Println(txResp)

	return &txResp.TxHash, nil
}

func main() {
	dataSourceCmd := flag.NewFlagSet("report", flag.ExitOnError)
	vesselImo := dataSourceCmd.String("imo", DefaultVesselIMO, "The IMO identifier of the ship you want to fetch data for")
	consolidateCmd := flag.NewFlagSet("consolidate", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'report' or 'consolidate' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "report", "r":
		dataSourceCmd.Parse(os.Args[2:])

		apiKey := os.Getenv("VESSEL_DATALASTIC_API_KEY")
		if apiKey == "" {
			fmt.Println("You need to specify the VESSEL_DATALASTIC_API_KEY environment variable to run this application.")
			return
		}

		fmt.Println("Fetching data for vessel with IMO " + *vesselImo)

		apiURL := "https://api.datalastic.com/api/v0/vessel_pro?api-key=" + apiKey + "&imo=" + *vesselImo // Replace with the actual API URL
		apiData, err := fetchVesselDataFromDatalastic(apiURL, *vesselImo)
		if err != nil {
			fmt.Println("Error fetching data: ", err)
			return
		}

		// 1. create data sets for vessels including outliers for ETA and departure port
		vesselDataSources, err := generateVesselData(apiData)
		for index, vesselData := range vesselDataSources {
			str, err := json.MarshalIndent(vesselData, "", "  ")
			if err != nil {
				fmt.Println("Error while marshalling: ", err)
				return
			}
			fmt.Printf("vessel data [%d]: %v\n", index, string(str))
		}

		// 2. create and submit transactions/data reports to the blockchain and wait for acceptance
		transactions, err := createAndSubmitDataReports(vesselDataSources)
		if err != nil {
			fmt.Println("Error while creating and submitting transactions: ", err)
			return
		}
		fmt.Println("Transactions: ", transactions)
	case "consolidate", "c":
		consolidateCmd.Parse(os.Args[2:])
		transaction, err := submitDataConsolidationRequest(*vesselImo)
		if err != nil {
			fmt.Println("Error while requesting data consolidation: ", err)
			return
		}
		fmt.Println("Transaction: ", transaction)
	default:
		fmt.Errorf("Invalid subcommand. Only supported subcommands are 'consolidate' and 'report' or 'c' and 'r'.")
		os.Exit(1)
	}
}

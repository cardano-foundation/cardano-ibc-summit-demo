package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VesselList:                 []Vessel{},
		ConsolidatedDataReportList: []ConsolidatedDataReport{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in vessel
	vesselIndexMap := make(map[string]struct{})

	for _, elem := range gs.VesselList {
		index := string(VesselKey(elem.Imo, elem.Ts, elem.Source))
		if _, ok := vesselIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vessel")
		}
		vesselIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in consolidatedDataReport
	consolidatedDataReportIndexMap := make(map[string]struct{})

	for _, elem := range gs.ConsolidatedDataReportList {
		index := string(ConsolidatedDataReportKey(elem.Imo, elem.Ts))
		if _, ok := consolidatedDataReportIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for consolidatedDataReport")
		}
		consolidatedDataReportIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

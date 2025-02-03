# Cardano Summit 2025 IBC Masterclass Demo
This application aims to demonstrate the Cardano's IBC connectivity capabilities by implementing a simple Oracle service running on a Cosmos SDK based chain connected to Cardano via IBC. It fetches information from public ship data APIs, uses Cosmos vote extensions for aligning on a value among different data sources, sends this data via IBC to Cardano where it is picked up by a smart contract that processes it and delivers an outcome back to the Cosmos chain via IBC.

## Description
The demonstrator uses 3 components
- a custom [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) based blockchain scaffolded with [Ignite](https://ignite.com/)
- a simple CLI tool written in Go that fetches tracking data for ships from [Datalastic](https://datalastic.com/)
- the [Cardano IBC implementation](https://github.com/cardano-foundation/cardano-ibc-incubator) in it's current form based on the [Mithril](https://github.com/input-output-hk/mithril)-based light client

### Cosmos chain
The Cosmos chain is super simple and just contains a single map to store the ship data. The data model used is:
```protobuf
message Vessel {
  string imo = 1; // index
  uint64 ts = 2; // index
  string source = 3; // index
  int32 lat = 4;
  int32 lon = 5;
  int32 speed = 6;
  int32 course = 7;
  int32 heading = 8;
  uint64 adt = 9;
  uint64 eta = 10;
  string name = 11;
  string destport = 12;
  string depport = 13;
  string mmsi = 14;
  string creator = 15;
}
```
The `vesseloracle` module consumes data of type Vessel, performance an outlier detection on received data reports and sends a consolidated data record via IBC to Cardano.

### The Go `datasource` process
The process fetches the aforementioned data from Datalastic via this endpoint
```
https://api.datalastic.com/api/v0/vessel_pro?api-key=API-KEY&uuid=b8625b67-7142-cfd1-7b85-595cebfe4191
```
The data is only fetched once but jitter is introduced and an outlier for the `eta` and `depport` fields to simulate a data source that is not functioning correctly.

Once the data reports are created messages are created and submitted as transactions to the Cosmos chain.

### Cardano
TBD but essentially picks up the data once landed via IBC on Cardano, calls an Aiken contract and submits the data back.

## How to run?
Clone the repo and enter the directory created
```bash
git clone git@github.com:cardano-foundation/cardano-ibc-summit-demo.git
cd cardano-ibc-summit-demo
```

Run the Cosmos chain by entering the chain directory and use ignite to start the chain
```bash
cd vesseloracle
ignite chain serve -r
```

TODO ... use Caribic to setup the Cardano chain and other IBC infrastructure

Execute the Go process by opening a new terminal and going into the root directory of the cloned repository and then
```bash
cd datasource
go mod tidy
go run . report -simulate
go run . consolidate
go run . transmit -ts $CONSOLIDATION_TIMESTAMP
```
The `CONSOLIDATEN_TIMESTAMP` is part of a log message in the main process of the sidechain after successfully submitting the consolidated data using the `go run . consolidate` command.

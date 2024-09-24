package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VesselList: []Vessel{},
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
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

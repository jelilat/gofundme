package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		GofundmeList: []Gofundme{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in gofundme
	gofundmeIdMap := make(map[uint64]bool)
	gofundmeCount := gs.GetGofundmeCount()
	for _, elem := range gs.GofundmeList {
		if _, ok := gofundmeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for gofundme")
		}
		if elem.Id >= gofundmeCount {
			return fmt.Errorf("gofundme id should be lower or equal than the last id")
		}
		gofundmeIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

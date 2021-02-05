package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/provenance-io/provenance/x/name/types"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyMinSegmentLength),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%d", GenMinSegmentLength(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyMaxSegmentLength),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%d", GenMaxSegmentLength(r))
			},
		),

		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyMaxNameLevels),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%d", GenMaxNameLevels(r))
			},
		),

		simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyAllowUnrestrictedNames),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%v", GenAllowUnrestrictedNames(r))
			},
		),
	}
}

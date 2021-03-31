package types

import (
	"fmt"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type DenomTestSuite struct {
	suite.Suite
}

func (s *DenomTestSuite) SetupTest() {}

func TestDenomTestSuite(t *testing.T) {
	suite.Run(t, new(DenomTestSuite))
}

type denomMetadataTestCase struct {
	name     string
	md       banktypes.Metadata
	wantInErr []string
}

func getValidateDenomMetadataTestCases() []denomMetadataTestCase {
	return []denomMetadataTestCase{
		{
			"base is not a valid coin denomination",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  nil,
				Base:        "x",
				Display:     "hash",
			},
			[]string{"denom metadata"},
		},
		{
			"display is not a valid coin denomination",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  nil,
				Base:        "hash",
				Display:     "x",
			},
			[]string{"denom metadata"},
		},
		{
			"first denom unit is not exponent 0",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "hash", Exponent: 1, Aliases: nil},
				},
				Base:        "hash",
				Display:     "hash",
			},
			[]string{"denom metadata"},
		},
		{
			"first denom unit is not base",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "hash",
				Display:     "hash",
			},
			[]string{"denom metadata"},
		},
		{
			"denom units not ordered",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"denom metadata"},
		},
		{
			"description too long",
			banktypes.Metadata{
				Description: strings.Repeat("d", maxDenomMetadataDescriptionLength+1),
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"description", fmt.Sprint(maxDenomMetadataDescriptionLength), fmt.Sprint(maxDenomMetadataDescriptionLength+1)},
		},
		{
			"no root coin name",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "hashx", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"root coin name"},
		},
		{
			"base prefix not SI",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "xhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "xhash",
				Display:     "hash",
			},
			[]string{"root coin name", "is not a SI prefix"},
		},
		{
			"alias duplicates other name",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: []string{"uhash"}},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"denom or alias", "is not unique", "uhash"},
		},
		{
			"denom duplicates other name",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
					{Denom: "nanohash", Exponent: 12, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"denom or alias", "is not unique", "nanohash"},
		},
		{
			"denom unit denom is not valid a coin denomination",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "x", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"denom metadata"},
		},
		{
			"denom unit denom exponent is incorrect",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 8, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"exponent", "hash", "0", "-9", "= 9", "8"},
		},
		{
			"denom unit alias is not valid a coin denomination",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: []string{strings.Repeat("x", 128)+"hash"}},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"invalid alias", "x"},
		},
		{
			"denom unit denom alias prefix mismatch",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
					{Denom: "megahash", Exponent: 15, Aliases: []string{"mhash"}},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			[]string{"SI prefix", "mhash", "megahash"},
		},
	}
}

func defaultTestParams() Params {
	return testParams(`^[a-zA-Z]{4,30}$`)
}

func testParams(regex string) Params {
	return Params{
		MaxTotalSupply:         0,
		EnableGovernance:       false,
		UnrestrictedDenomRegex: regex,
	}
}

func (s *DenomTestSuite) TestValidateDenomMetadataBasic() {
	tests := getValidateDenomMetadataTestCases()

	for _, tc := range tests {
		s.T().Run(tc.name, func(t *testing.T) {
			err := ValidateDenomMetadataBasic(tc.md)
			if len(tc.wantInErr) > 0 {
				require.Error(t, err, "ValidateDenomMetadataBasic expected error")
				for _, e := range tc.wantInErr {
					assert.Contains(t, err.Error(), e, "ValidateDenomMetadataBasic expected in error message")
				}
			} else {
				require.NoError(t, err, "ValidateDenomMetadataBasic unexpected error")
			}
		})
	}
}

func (s *DenomTestSuite) TestValidateDenomMetadataExtended() {
	basicTests := getValidateDenomMetadataTestCases()

	// Should call ValidateDenomMetadataBasic, so all these should apply here too.
	for _, tc := range basicTests {
		s.T().Run("basic " + tc.name, func(t *testing.T) {
			err := ValidateDenomMetadataExtended(tc.md, nil, StatusProposed, defaultTestParams())
			if len(tc.wantInErr) > 0 {
				require.Error(t, err, "ValidateDenomMetadataExtended expected error")
				for _, e := range tc.wantInErr {
					assert.Contains(t, err.Error(), e, "ValidateDenomMetadataExtended expected in error message")
				}
			} else {
				require.NoError(t, err, "ValidateDenomMetadataExtended unexpected error")
			}
		})
	}

	tests := []struct {
		name string
		proposed     banktypes.Metadata
		existing     *banktypes.Metadata
		markerStatus MarkerStatus
		params       Params
		wantInErr    []string
	}{
		{
			"marker status undefined",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			nil,
			StatusUndefined,
			defaultTestParams(),
			[]string{"cannot add or update denom metadata", "undefined"},
		},
		{
			"marker status destroyed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			nil,
			StatusDestroyed,
			defaultTestParams(),
			[]string{"cannot add or update denom metadata", "destroyed"},
		},
		{
			"marker status cancelled",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			nil,
			StatusCancelled,
			defaultTestParams(),
			[]string{"cannot add or update denom metadata", "cancelled"},
		},
		{
			"denom fails extra regex",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			nil,
			StatusProposed,
			testParams(`^[nu]hash$`),
			[]string{"fails unrestricted marker denom regex", "hash"},
		},
		{
			"alias fails extra regex",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			nil,
			StatusProposed,
			testParams(`^[nu]?hash$`),
			[]string{"fails unrestricted marker denom regex", "nanohash"},
		},
		{
			"invalid unrestricted marker denom regex",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			nil,
			StatusProposed,
			testParams(`(foo`),
			[]string{"error parsing regexp"},
		},
		{
			"base changed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "uhash",
				Display:     "hash",
			},
			StatusProposed,
			defaultTestParams(),
			[]string{"denom metadata base value cannot be changed"},
		},
		{
			"active denom unit removed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusActive,
			defaultTestParams(),
			[]string{"cannot remove denom unit", "uhash"},
		},
		{
			"finalized denom unit removed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusFinalized,
			defaultTestParams(),
			[]string{"cannot remove denom unit", "uhash"},
		},
		{
			"proposed denom unit removed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusProposed,
			defaultTestParams(),
			[]string{},
		},
		{
			"active denom unit denom changed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "microhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusActive,
			defaultTestParams(),
			[]string{"denom unit Denom", "uhash", "microhash"},
		},
		{
			"finalized denom unit denom changed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "microhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusFinalized,
			defaultTestParams(),
			[]string{"denom unit Denom", "uhash", "microhash"},
		},
		{
			"proposed denom unit denom changed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "microhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusProposed,
			defaultTestParams(),
			[]string{},
		},
		{
			"active denom unit alias removed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusActive,
			defaultTestParams(),
			[]string{"cannot remove alias", "nanohash", "nhash"},
		},
		{
			"finalized denom unit alias removed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusFinalized,
			defaultTestParams(),
			[]string{"cannot remove alias", "nanohash", "nhash"},
		},
		{
			"proposed denom unit alias removed",
			banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: nil},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			&banktypes.Metadata{
				Description: "a description",
				DenomUnits:  []*banktypes.DenomUnit{
					{Denom: "nhash", Exponent: 0, Aliases: []string{"nanohash"}},
					{Denom: "uhash", Exponent: 3, Aliases: nil},
					{Denom: "hash", Exponent: 9, Aliases: nil},
				},
				Base:        "nhash",
				Display:     "hash",
			},
			StatusProposed,
			defaultTestParams(),
			[]string{},
		},
	}

	for _, tc := range tests {
		s.T().Run(tc.name, func(t *testing.T) {
			err := ValidateDenomMetadataExtended(tc.proposed, tc.existing, tc.markerStatus, tc.params)
			if len(tc.wantInErr) > 0 {
				require.Error(t, err, "ValidateDenomMetadataExtended expected error")
				for _, e := range tc.wantInErr {
					assert.Contains(t, err.Error(), e, "ValidateDenomMetadataExtended expected in error message")
				}
			} else {
				require.NoError(t, err, "ValidateDenomMetadataExtended unexpected error")
			}
		})
	}
}

func (s *DenomTestSuite) TestGetRootCoinName() {
	tests := []struct {
		name     string
		md       banktypes.Metadata
		expected string
	}{
		{
			"empty metadata",
			banktypes.Metadata{},
			"",
		},
		{
			"only one name",
			banktypes.Metadata{
				DenomUnits:  []*banktypes.DenomUnit{
					{
						Denom:    "onename",
						Aliases:  nil,
					},
				},
			},
			"",
		},
		{
			"no common root",
			banktypes.Metadata{
				DenomUnits:  []*banktypes.DenomUnit{
					{
						Denom:    "onename",
						Aliases:  []string{"another"},
					},
				},
			},
			"",
		},
		{
			"simple test",
			banktypes.Metadata{
				DenomUnits:  []*banktypes.DenomUnit{
					{
						Denom:    "onename",
						Aliases:  []string{"twoname"},
					},
				},
			},
			"name",
		},
		{
			"real-use test",
			banktypes.Metadata{
				DenomUnits:  []*banktypes.DenomUnit{
					{
						Denom:    "nanohash",
						Aliases:  []string{"nhash"},
					},
					{
						Denom:    "hash",
						Aliases:  nil,
					},
					{
						Denom:    "kilohash",
						Aliases:  []string{"khash"},
					},
				},
			},
			"hash",
		},
	}

	for _, tc := range tests {
		s.T().Run(tc.name, func(t *testing.T) {
			actual := GetRootCoinName(tc.md)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
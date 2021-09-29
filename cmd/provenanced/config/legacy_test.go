package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	v34config "github.com/provenance-io/provenance/cmd/provenanced/config/legacy/tendermint_0_34/config"
	v35config "github.com/provenance-io/provenance/cmd/provenanced/config/legacy/tendermint_0_35/config"
)

type LegacyTestSuite struct {
	suite.Suite

	Home string
}

func TestLegacyTestSuite(t *testing.T) {
	suite.Run(t, new(LegacyTestSuite))
}

func (s *LegacyTestSuite) SetupTest() {
	s.Home = s.T().TempDir()
	s.T().Logf("%s Home: %s", s.T().Name(), s.Home)
}

func containsString(a []string, s string) bool {
	for _, t := range a {
		if s == t {
			return true
		}
	}
	return false
}

func (s *LegacyTestSuite) TestCompare34and35() {
	v34 := v34config.DefaultConfig()
	v35 := v35config.DefaultConfig()

	v34Map := MakeFieldValueMap(v34, true)
	v35Map := MakeFieldValueMap(v35, true)

	unchanged := []string{}
	added := []string{}
	removed := []string{}
	toDashes := []string{}

	for key34 := range v34Map {
		if _, ok := v35Map[key34]; ok {
			unchanged = append(unchanged, key34)
			continue
		}
		if _, ok := v35Map[strings.Replace(key34, "_", "-", -1)]; ok {
			toDashes = append(toDashes, key34)
		} else {
			removed = append(removed, key34)
		}
	}

	for key35 := range v35Map {
		if _, ok := v34Map[key35]; ok {
			continue
		}
		if _, ok := v34Map[strings.Replace(key35, "-", "_", -1)]; ok {
			continue
		}
		added = append(added, key35)
	}

	sortKeys(unchanged)
	sortKeys(added)
	sortKeys(removed)
	sortKeys(toDashes)

	printStrings := func(header string, vals []string) {
		fmt.Printf("%s:\n", header)
		for _, val := range vals {
			fmt.Printf("  %s\n", val)
		}
		fmt.Printf("\n")
	}

	printStrings("unchanged", unchanged)
	printStrings("added", added)
	printStrings("removed", removed)
	printStrings("toDashes", toDashes)
	s.T().Fail()
}

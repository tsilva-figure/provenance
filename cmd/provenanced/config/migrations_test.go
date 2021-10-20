package config

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigMigrationsTestSuite struct {
	suite.Suite

	// Home is a temp directory that can be used to store files for a test.
	// It is different for each test function.
	Home string
}

func TestConfigMigrationsTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigMigrationsTestSuite))
}

func (s *ConfigMigrationsTestSuite) SetupTest() {
	s.Home = s.T().TempDir()
	s.T().Logf("%s Home: %s", s.T().Name(), s.Home)
}

func (s *ConfigMigrationsTestSuite) TestUniqueKeyEntries() {
	// This test makes sure that a key is only listed once among the
	// addedKeys, removedKeys, toDashes, and changedKeys variables.
	keySources := make(map[string][]string)
	addKey := func(key, source string) {
		keySources[key] = append(keySources[key], source)
	}
	for _, key := range addedKeys {
		addKey(key, "addedKeys")
	}
	for _, key := range removedKeys {
		addKey(key, "removedKeys")
	}
	for _, key := range toDashesKeys {
		addKey(key, "toDashesKeys")
	}
	for oldKey, newKey := range changedKeys {
		addKey(oldKey, "changedKeys-old")
		addKey(newKey, "changedKeys-new")
	}
	for key, sources := range keySources {
		s.Assert().Len(sources, 1, key)
	}
}

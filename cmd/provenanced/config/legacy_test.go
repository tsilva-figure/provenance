package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
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

func stringSliceContains(ss []string, s string) bool {
	for _, t := range ss {
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

func (s *LegacyTestSuite) TestCompareConfigToFileEntries() {
	// Here's how I got these:
	// Created the file where I could get at it using this line in a test:
	// v35config.WriteConfigFile("/Users/danielwedul/random-work/prov-configs/tm35", v35config.DefaultConfig())
	// In a terminal:
	//   cd /Users/danielwedul/random-work/prov-configs/tm35/config
	//   toml-to-json -p config.toml > config.json
	//   json_info -r -f config.json --just-paths | sed 's/^\.\["//; s/^\.//; s/\["/./g; s/"\]//g;' | grep -vF '[' | pbcopy
	// Paste it in here.
	// Reorder the base entries to the top and remove the lines that have just the group name.
	fileEntries := strings.Split(`abci
db-backend
db-dir
filter-peers
genesis-file
log-format
log-level
mode
moniker
node-key-file
proxy-app
blocksync.enable
blocksync.version
consensus.create-empty-blocks
consensus.create-empty-blocks-interval
consensus.double-sign-check-height
consensus.peer-gossip-sleep-duration
consensus.peer-query-maj23-sleep-duration
consensus.skip-timeout-commit
consensus.timeout-commit
consensus.timeout-precommit
consensus.timeout-precommit-delta
consensus.timeout-prevote
consensus.timeout-prevote-delta
consensus.timeout-propose
consensus.timeout-propose-delta
consensus.wal-file
instrumentation.max-open-connections
instrumentation.namespace
instrumentation.prometheus
instrumentation.prometheus-listen-addr
mempool.broadcast
mempool.cache-size
mempool.keep-invalid-txs-in-cache
mempool.max-batch-bytes
mempool.max-tx-bytes
mempool.max-txs-bytes
mempool.recheck
mempool.size
mempool.ttl-duration
mempool.ttl-num-blocks
mempool.version
p2p.addr-book-file
p2p.addr-book-strict
p2p.allow-duplicate-ip
p2p.bootstrap-peers
p2p.dial-timeout
p2p.external-address
p2p.flush-throttle-timeout
p2p.handshake-timeout
p2p.laddr
p2p.max-connections
p2p.max-incoming-connection-attempts
p2p.max-num-inbound-peers
p2p.max-num-outbound-peers
p2p.max-packet-msg-payload-size
p2p.persistent-peers
p2p.persistent-peers-max-dial-period
p2p.pex
p2p.private-peer-ids
p2p.queue-type
p2p.recv-rate
p2p.seeds
p2p.send-rate
p2p.unconditional-peer-ids
p2p.upnp
p2p.use-legacy
priv-validator.certificate-authority
priv-validator.client-certificate-file
priv-validator.key-file
priv-validator.laddr
priv-validator.state-file
priv-validator.validator-client-key-file
rpc.cors-allowed-headers
rpc.cors-allowed-methods
rpc.cors-allowed-origins
rpc.grpc-laddr
rpc.grpc-max-open-connections
rpc.laddr
rpc.max-body-bytes
rpc.max-header-bytes
rpc.max-open-connections
rpc.max-subscription-clients
rpc.max-subscriptions-per-client
rpc.pprof-laddr
rpc.timeout-broadcast-tx-commit
rpc.tls-cert-file
rpc.tls-key-file
rpc.unsafe
statesync.chunk-request-timeout
statesync.discovery-time
statesync.enable
statesync.fetchers
statesync.rpc-servers
statesync.temp-dir
statesync.trust-hash
statesync.trust-height
statesync.trust-period
statesync.use-p2p
tx-index.indexer
tx-index.psql-conn`, "\n")

	v35Cfg := v35config.DefaultConfig()
	v35Map := removeUndesirableTmConfigEntries(MakeFieldValueMap(v35Cfg, true))

	configEntries := make([]string, 0)
	inConfigButNotFile := make([]string, 0)
	for v35key := range v35Map {
		configEntries = append(configEntries, v35key)
		if !stringSliceContains(fileEntries, v35key) {
			inConfigButNotFile = append(inConfigButNotFile, v35key)
		}
	}
	sortKeys(inConfigButNotFile)

	inFileButNotConfig := make([]string, 0)
	for _, fileKey := range fileEntries {
		if !stringSliceContains(configEntries, fileKey) {
			inFileButNotConfig = append(inFileButNotConfig, fileKey)
		}
	}
	sortKeys(inFileButNotConfig)

	s.Assert().Len(inConfigButNotFile, 0, "In config, but not file.")
	s.Assert().Len(inFileButNotConfig, 0, "In file, but not config.")
}

func (s *LegacyTestSuite) TestRead34FileWith35Struct() {
	v34 := v34config.DefaultConfig()
	confFile := filepath.Join(s.Home, "config.toml")
	v34config.WriteConfigFile(confFile, v34)

	vpr := viper.New()
	vpr.SetConfigFile(confFile)
	err := vpr.ReadInConfig()
	s.Require().NoError(err, "reading config into viper")

	v35 := v35config.DefaultConfig()
	err = vpr.Unmarshal(v35)
	s.Require().NoError(err, "unmarshaling conf from viper")

	otherKeys := make([]string, 0, len(v35.Other))
	for key := range v35.Other {
		otherKeys = append(otherKeys, key)
	}
	sortKeys(otherKeys)
	for _, key := range otherKeys {
		val := v35.Other[key]
		fmt.Printf("%s: %#v\n", key, val)
	}
	s.Assert().Len(otherKeys, 0, "other keys")
}

func (s *LegacyTestSuite) TestRead34FileWithMap() {
	v34 := v34config.DefaultConfig()
	confFile := filepath.Join(s.Home, "config.toml")
	v34config.WriteConfigFile(confFile, v34)

	vpr := viper.New()
	vpr.SetConfigFile(confFile)
	err := vpr.ReadInConfig()
	s.Require().NoError(err, "reading config into viper")

	v35 := map[string]interface{}{}
	err = vpr.Unmarshal(&v35)
	s.Require().NoError(err, "unmarshaling conf from viper")

	printMap := func(header string, m map[string]interface{}) []string {
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		sortKeys(keys)
		fmt.Printf("%s:\n", header)
		for _, key := range keys {
			fmt.Printf("%s: %#v\n", key, m[key])
		}
		return keys
	}
	printMapStr := func(header string, m map[string]string) []string {
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		sortKeys(keys)
		fmt.Printf("%s:\n", header)
		for _, key := range keys {
			fmt.Printf("%s: \"%s\"\n", key, m[key])
		}
		return keys
	}

	v35Keys := printMap("base", v35)
	s.Assert().Len(v35Keys, 0, "base keys")

	v35Consensus := v35["consensus"].(map[string]interface{})
	v35ConsensusKeys := printMap("consensus", v35Consensus)
	s.Assert().Len(v35ConsensusKeys, 0, "consensus keys")

	v35Flat := flattenMap(v35)
	printMapStr("flattened", v35Flat)
}

func flattenMap(m map[string]interface{}) map[string]string {
	rv := map[string]string{}
	for key, val := range m {
		if valm, ok := val.(map[string]interface{}); ok {
			for subkey, subval := range flattenMap(valm) {
				rv[key+"."+subkey] = subval
			}
		} else {
			rv[key] = fmt.Sprintf("%v", val)
		}
	}
	return rv
}

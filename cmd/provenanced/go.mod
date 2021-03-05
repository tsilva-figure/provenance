module github.com/provenance-io/provenance/cmd/provenanced

go 1.15

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

require (
	github.com/cosmos/cosmos-sdk v0.41.4
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/provenance-io/provenance v0.1.10
	github.com/rs/zerolog v1.20.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
)

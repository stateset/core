module github.com/stateset/core

go 1.24

require (
	cosmossdk.io/api v0.9.2
	cosmossdk.io/math v1.5.3
	cosmossdk.io/core v1.0.0
	cosmossdk.io/depinject v1.2.1
	cosmossdk.io/errors v1.0.2
	cosmossdk.io/log v1.6.0
	cosmossdk.io/x/evidence v0.1.1
	cosmossdk.io/x/feegrant v0.1.1
	cosmossdk.io/x/upgrade v0.1.4
	cosmossdk.io/store v1.1.1
	github.com/CosmWasm/wasmd v0.50.0
	github.com/cometbft/cometbft v0.38.12
	github.com/cometbft/cometbft-db v0.14.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.5
	github.com/cosmos/cosmos-sdk v0.50.10
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/gogoproto v1.7.0
	github.com/cosmos/ibc-go/v8 v8.5.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.4
	github.com/gorilla/mux v1.8.1
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/prometheus/client_golang v1.20.5
	github.com/spf13/cast v1.7.0
	github.com/spf13/cobra v1.8.1
	github.com/stretchr/testify v1.9.0
	google.golang.org/genproto/googleapis/api v0.0.0-20241210142825-c6dcf9bb7e75
	google.golang.org/grpc v1.69.2
	gopkg.in/yaml.v2 v2.4.0
)



replace cosmossdk.io/store => github.com/cosmos/cosmos-sdk/store v1.1.2

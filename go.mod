module github.com/calvinlauco/cosmostxdecoder

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/gogo/protobuf v1.3.1
	github.com/stretchr/testify v1.6.1
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.40.0-rc3

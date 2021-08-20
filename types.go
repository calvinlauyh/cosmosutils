package cosmosutils

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctypes "github.com/cosmos/ibc-go/modules/core/types"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	nfttypes "github.com/crypto-org-chain/chain-main/v3/x/nft/types"
)

func RegisterDefaultInterfaces(interfaceRegistry types.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)

	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	proposal.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	ibctypes.RegisterInterfaces(interfaceRegistry)
	ibctransfertypes.RegisterInterfaces(interfaceRegistry)
	nfttypes.RegisterInterfaces(interfaceRegistry)
}

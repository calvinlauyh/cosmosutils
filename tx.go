package cosmostxdecoder

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type Tx struct {
	cosmostypes.Tx

	codec codec.ProtoCodecMarshaler
}

func (tx *Tx) MarshalToJSON() ([]byte, error) {
	return authtx.DefaultJSONTxEncoder(tx.codec)(tx.Tx)
}

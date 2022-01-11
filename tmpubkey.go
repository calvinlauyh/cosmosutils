package cosmosutils

import (
	"fmt"

	"github.com/cosmos/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func ConsensusNodeAddressFromTmPubKey(bech32Prefix string, pubKey []byte) (string, error) {
	cosmosPubKey := &ed25519.PubKey{
		Key: pubKey,
	}

	conv, err := bech32.ConvertBits(cosmosPubKey.Address().Bytes(), 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("error converting tendermint public key to bech32 bits: %v", err)
	}
	address, err := bech32.Encode(bech32Prefix, conv)
	if err != nil {
		return "", fmt.Errorf("error encoding tendermint public key bits to consensus address: %v", err)
	}

	return address, nil
}

func ConsensusNodeAddressFromPubKey(bech32Prefix string, consensusNodePubKey string) (string, error) {
	_, conv, err := bech32.Decode(consensusNodePubKey, 1023)
	if err != nil {
		return "", fmt.Errorf("error converting consensus node pubkey to address")
	}

	pkToUnmarshal, err := bech32.ConvertBits(conv, 5, 8, false)
	if err != nil {
		return "", fmt.Errorf("error converting bech32 bits to tendermint public key: %v", err)
	}
	var pubKey cryptotypes.PubKey
	legacy.Cdc.MustUnmarshal(pkToUnmarshal, &pubKey)

	conv, err = bech32.ConvertBits(pubKey.Address().Bytes(), 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("error converting tendermint public key to bech32 bits: %v", err)
	}
	address, err := bech32.Encode(bech32Prefix, conv)
	if err != nil {
		return "", fmt.Errorf("error encoding tendermint public key bits to consensus address: %v", err)
	}

	return address, nil
}

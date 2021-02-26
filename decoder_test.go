package cosmostxdecoder_test

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/calvinlauco/cosmostxdecoder"
)

func TestDecodeInvalidBase64Tx(t *testing.T) {
	decoder := cosmostxdecoder.NewDecoder()
	_, err := decoder.DecodeBase64("!@#$%^&*()_+=-`")

	assert.EqualError(t, err, "illegal base64 data at input byte 0")
}

func TestDecodeInvalidTxBytes(t *testing.T) {
	decoder := cosmostxdecoder.NewDecoder()
	_, err := decoder.Decode([]byte("INVALID"))

	assert.EqualError(t, err, "errUnknownField \"*tx.TxRaw\": {TagNum: 9, WireType:\"fixed64\"}: tx parse error")
}

func TestDecodeTxBytes(t *testing.T) {
	anyTxBytes, _ := hex.DecodeString("0A94010A91010A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E6412710A2B7463726F31666D70726D30736A79366C7A396C6C7637726C746E307632617A7A7763777A766B326C73796E122B7463726F31373832676E39687A7161766563756B64617171636C76736E70636B346D747A3376777A70786C1A150A08626173657463726F120931303030303030303012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210359A154BA210C489DA46626A4D631C6F8A471A2FBDA342164DD5FC4A158901F2612040A02087F180412150A100A08626173657463726F12043130303010904E1A40E812FBA1D50115CDDD534C7675B838E6CE7169CDD5AD312182696CABB76F05B82B9557EE1A2842A5579B363F0A5F2E487F7228A81FBF81E125E58F264A29DA07")
	expected := "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\",\"to_address\":\"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl\",\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"100000000\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[{\"public_key\":{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m\"},\"mode_info\":{\"single\":{\"mode\":\"SIGN_MODE_LEGACY_AMINO_JSON\"}},\"sequence\":\"4\"}],\"fee\":{\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"1000\"}],\"gas_limit\":\"10000\",\"payer\":\"\",\"granter\":\"\"}},\"signatures\":[\"6BL7odUBFc3dU0x2dbg45s5xac3VrTEhgmlsq7dvBbgrlVfuGihCpVebNj8KXy5If3IoqB+/geEl5Y8mSinaBw==\"]}"
	decoder := cosmostxdecoder.DefaultDecoder

	tx, err := decoder.Decode(anyTxBytes)
	assert.NoError(t, err)
	fmt.Println(tx)

	actual, err := tx.MarshalToJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestDecodeAminoJSONSignedTx(t *testing.T) {
	anyAminoSignModeMsgSendTx := "CpQBCpEBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnEKK3Rjcm8xZm1wcm0wc2p5Nmx6OWxsdjdybHRuMHYyYXp6d2N3enZrMmxzeW4SK3Rjcm8xNzgyZ245aHpxYXZlY3VrZGFxcWNsdnNucGNrNG10ejN2d3pweGwaFQoIYmFzZXRjcm8SCTEwMDAwMDAwMBJpClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDWaFUuiEMSJ2kZiak1jHG+KRxovvaNCFk3V/EoViQHyYSBAoCCH8YBBIVChAKCGJhc2V0Y3JvEgQxMDAwEJBOGkDoEvuh1QEVzd1TTHZ1uDjmznFpzdWtMSGCaWyrt28FuCuVV+4aKEKlV5s2PwpfLkh/ciioH7+B4SXljyZKKdoH"
	expected := "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\",\"to_address\":\"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl\",\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"100000000\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[{\"public_key\":{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m\"},\"mode_info\":{\"single\":{\"mode\":\"SIGN_MODE_LEGACY_AMINO_JSON\"}},\"sequence\":\"4\"}],\"fee\":{\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"1000\"}],\"gas_limit\":\"10000\",\"payer\":\"\",\"granter\":\"\"}},\"signatures\":[\"6BL7odUBFc3dU0x2dbg45s5xac3VrTEhgmlsq7dvBbgrlVfuGihCpVebNj8KXy5If3IoqB+/geEl5Y8mSinaBw==\"]}"
	decoder := cosmostxdecoder.DefaultDecoder

	tx, err := decoder.DecodeBase64(anyAminoSignModeMsgSendTx)
	assert.NoError(t, err)

	actual, err := tx.MarshalToJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestDecodeDirectSignedTx(t *testing.T) {
	anyDirectSignedTx := "CpQBCpEBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnEKK3Rjcm8xZm1wcm0wc2p5Nmx6OWxsdjdybHRuMHYyYXp6d2N3enZrMmxzeW4SK3Rjcm8xNzgyZ245aHpxYXZlY3VrZGFxcWNsdnNucGNrNG10ejN2d3pweGwaFQoIYmFzZXRjcm8SCTEwMDAwMDAwMBJpClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDWaFUuiEMSJ2kZiak1jHG+KRxovvaNCFk3V/EoViQHyYSBAoCCAEYBBIVChAKCGJhc2V0Y3JvEgQxMDAwEJBOGkAKlWtAxcXT06uPxtdgT2OthBz3WaW4Kz0fJ3XVq3VLxyIqL12OkpbTmKJUsK8avHzu4pz1CXYBlaUoNiPZvqUr"
	expected := "{\"body\":{\"messages\":[{\"@type\":\"/cosmos.bank.v1beta1.MsgSend\",\"from_address\":\"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn\",\"to_address\":\"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl\",\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"100000000\"}]}],\"memo\":\"\",\"timeout_height\":\"0\",\"extension_options\":[],\"non_critical_extension_options\":[]},\"auth_info\":{\"signer_infos\":[{\"public_key\":{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m\"},\"mode_info\":{\"single\":{\"mode\":\"SIGN_MODE_DIRECT\"}},\"sequence\":\"4\"}],\"fee\":{\"amount\":[{\"denom\":\"basetcro\",\"amount\":\"1000\"}],\"gas_limit\":\"10000\",\"payer\":\"\",\"granter\":\"\"}},\"signatures\":[\"CpVrQMXF09Orj8bXYE9jrYQc91mluCs9Hyd11at1S8ciKi9djpKW05iiVLCvGrx87uKc9Ql2AZWlKDYj2b6lKw==\"]}"
	decoder := cosmostxdecoder.DefaultDecoder

	tx, err := decoder.DecodeBase64(anyDirectSignedTx)
	assert.NoError(t, err)

	actual, err := tx.MarshalToJSON()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestDecodeIBCBase64Tx(t *testing.T) {
	decoder := cosmostxdecoder.DefaultDecoder
	//https://www.mintscan.io/cosmos/txs/E3944CB1AF60EB1649B2DC9EBB6C67FFCCE17C6E44259417F4C4A657D9206B6E
	_, err := decoder.DecodeBase64("CqABCpEBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnEKLWNvc21vczF6ZHNmZTN0c3dza2h4eWh1amNuNGFuODdxcGZ1OWdmNG5renJlbhItY29zbW9zMWhwMmc1M3N0cWs3cnBsd3N4ZzA2MHMwNXRmdHpqNGRmZDNzYTB5GhEKBXVhdG9tEggyMzc5MTgxORIKMTU5NjQyNjI0NRJlCk4KRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDbihAoG4eoUY0R8iAGJ1uQEJKb8XwAatZR0f+HzrnFlUSBAoCCH8SEwoNCgV1YXRvbRIEMjAwMBCA8QQaQJ0OaKsjTwhwtJVwXemWhJR2etmR4e/z+o+g2PMfo1J/J/9R/qgLnxVjbRfAPjMf6XfLQcxdlflR7DP+CR+5yA4=")
	assert.NoError(t, err)
}

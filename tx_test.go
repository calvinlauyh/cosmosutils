package cosmostxdecoder_test

import (
    "bytes"
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "testing"

    "github.com/calvinlauco/cosmostxdecoder"
)

func TestCosmosTxUnmarshal(t *testing.T) {
    anyMsgSendTx := "CpQBCpEBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEnEKK3Rjcm8xZm1wcm0wc2p5Nmx6OWxsdjdybHRuMHYyYXp6d2N3enZrMmxzeW4SK3Rjcm8xNzgyZ245aHpxYXZlY3VrZGFxcWNsdnNucGNrNG10ejN2d3pweGwaFQoIYmFzZXRjcm8SCTEwMDAwMDAwMBJpClAKRgofL2Nvc21vcy5jcnlwdG8uc2VjcDI1NmsxLlB1YktleRIjCiEDWaFUuiEMSJ2kZiak1jHG+KRxovvaNCFk3V/EoViQHyYSBAoCCH8YBBIVChAKCGJhc2V0Y3JvEgQxMDAwEJBOGkDoEvuh1QEVzd1TTHZ1uDjmznFpzdWtMSGCaWyrt28FuCuVV+4aKEKlV5s2PwpfLkh/ciioH7+B4SXljyZKKdoH"
    decoder := cosmostxdecoder.DefaultDecoder

    tx, err := decoder.DecodeBase64(anyMsgSendTx)
    assert.NoError(t, err)

    jsonBytes, err := tx.MarshalToJSON()
    assert.NoError(t, err)

    var cosmosTx *cosmostxdecoder.CosmosTx
    jsonDecoder := json.NewDecoder(bytes.NewReader(jsonBytes))
    jsonDecoder.DisallowUnknownFields()
    jsonDecodeErr := jsonDecoder.Decode(&cosmosTx)
    assert.NoError(t, jsonDecodeErr)
}
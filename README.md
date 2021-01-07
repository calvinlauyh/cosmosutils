# Golang Cosmos SDK (Stargate) Transaction Decoder

[![GoDoc](https://godoc.org/github.com/calvinlauco/cosmostxdecoder?status.svg)](http://godoc.org/github.com/calvinlauco/cosmostxdecoder)
[![license](https://img.shields.io/github/license/calvinlauco/cosmostxdecoder.svg)](https://github.com/calvinlauco/cosmostxdecoder/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/calvinlauco/cosmostxdecoder/branch/master/graph/badge.svg)](https://codecov.io/gh/calvinlauco/cosmostxdecoder)
[![GitHub Actions](https://github.com/calvinlauco/cosmostxdecoder/workflows/build/badge.svg?branch=master)](https://github.com/calvinlauco/cosmostxdecoder/actions)

A Golang library and tool to decode Cosmos SDK (Stargate) transaction in base64 and hex encoded format to JSON.

## Pre-requisite

[Golang](https://golang.org/dl/)

## CLI

The project shipped with two binaries:
workflows
- `decode-cosmosbase64tx [base64-encoded-tx]`
    - Decode transaction encoded in base64. This is usually find in Cosmos SDK CLI and Tendermint block API endpoint.
- `decode-cosmostx [hex-encoded-tx]`
    - Decode transaction encoded in hex string
- `decode-cosmospubkey [pubkey] -p [prefix][-t]`
    - Convert bech32 consensus pubkey/ tenedermint pubkey to bech32 consensus address

### Build from source

```baseh
make all
```

### Decode your transaction

```bash
$ ./build/decode-cosmosbase64tx CowGCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTkyM3B6MDNtaGphenRnY3YzZ2V5MGhqMGFtd3gwMmR5c2thdTUyCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMXh3ZDNrOHh0ZXJkZWZ0M254cWc5MnN6aHB6NnZ4NDNxc3BkcHc2CpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTh5bGNoZ214eXBodzNjdHNsNzVuNTN1amVxdWttbWFnMm42eDNmCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZzOHI2enhtcjVuYzg2ajhjcGNtam1jY2Y4czJjYWZ4enQ1YWxxCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZqYTVuc3h6N2dzcXc0emNjdXV5OHI3cGpuam1jN2RzY2RsMnZ6EmwKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQL6Inwmwwd0nDUwtu9S8U0E+TU86f92eeo/ZUJfq+O1tRIECgIIARieHRIXChEKCGJhc2V0Y3JvEgUxNTAwMBDwkwkaQDfjeHPhtkdkDG3JyqECSN2DTIZeTC3Z2dK82HL1qshIH6dvMvT2JP4NGhmcQW/JK97sZ+FMdxe98GJxQNLfZfk=
{"body":{"messages":[{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn","to_address":"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl","amount":[{"denom":"basetcro","amount":"100000000"}]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[{"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m"},"mode_info":{"single":{"mode":"SIGN_MODE_LEGACY_AMINO_JSON"}},"sequence":"4"}],"fee":{"amount":[{"denom":"basetcro","amount":"1000"}],"gas_limit":"10000","payer":"","granter":""}},"signatures":["6BL7odUBFc3dU0x2dbg45s5xac3VrTEhgmlsq7dvBbgrlVfuGihCpVebNj8KXy5If3IoqB+/geEl5Y8mSinaBw=="]}

$ ./build/decode-cosmostx 0A94010A91010A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E6412710A2B7463726F31666D70726D30736A79366C7A396C6C7637726C746E307632617A7A7763777A766B326C73796E122B7463726F31373832676E39687A7161766563756B64617171636C76736E70636B346D747A3376777A70786C1A150A08626173657463726F120931303030303030303012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210359A154BA210C489DA46626A4D631C6F8A471A2FBDA342164DD5FC4A158901F2612040A02087F180412150A100A08626173657463726F12043130303010904E1A40E812FBA1D50115CDDD534C7675B838E6CE7169CDD5AD312182696CABB76F05B82B9557EE1A2842A5579B363F0A5F2E487F7228A81FBF81E125E58F264A29DA07
{"body":{"messages":[{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn","to_address":"tcro1782gn9hzqavecukdaqqclvsnpck4mtz3vwzpxl","amount":[{"denom":"basetcro","amount":"100000000"}]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[{"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A1mhVLohDEidpGYmpNYxxvikcaL72jQhZN1fxKFYkB8m"},"mode_info":{"single":{"mode":"SIGN_MODE_LEGACY_AMINO_JSON"}},"sequence":"4"}],"fee":{"amount":[{"denom":"basetcro","amount":"1000"}],"gas_limit":"10000","payer":"","granter":""}},"signatures":["6BL7odUBFc3dU0x2dbg45s5xac3VrTEhgmlsq7dvBbgrlVfuGihCpVebNj8KXy5If3IoqB+/geEl5Y8mSinaBw=="]}
```

### Convert tendermint pubkey or bech32 consensus pubkey to bech32 consensus address
```bash
$ ./build/decode-cosmospubkey Og8ZfQTHFgTBGD5qoyo5NpyJCJRddC+WuSPtyZtlE7E= -t 
Bech32 Validator Consensus Address: tcrocnclcons1wqajdt4qseasx4e8r8fz7juwdkfu4quvt9e87u
$ ./build/decode-cosmospubkey tcrocnclconspub1zcjduepq8g83jlgycutqfsgc8e42x23ex6wgjzy5t46zl94ey0kunxm9zwcsuzkxpr
Bech32 Validator Consensus Address: tcrocnclcons1wqajdt4qseasx4e8r8fz7juwdkfu4quvt9e87u
```
## Library

You can also use this package as Golang library

### `DefaultDecoder`

The `DefaultDecoder` is a pre-created decoder and has all the Cosmos built-in modules interfaces registered. You can use it right away:

```go
package yourpackage

import (
    "fmt"

    "github.com/calvinlauco/cosmostxdecoder"
)

func DecodeTxBytes() {
    decoder := cosmostxdecoder.DefaultDecoder
    anyTxBytes := []byte("0A94010A91010A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E6412710A2B7463726F31666D70726D30736A79366C7A396C6C7637726C746E307632617A7A7763777A766B326C73796E122B7463726F31373832676E39687A7161766563756B64617171636C76736E70636B346D747A3376777A70786C1A150A08626173657463726F120931303030303030303012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210359A154BA210C489DA46626A4D631C6F8A471A2FBDA342164DD5FC4A158901F2612040A02087F180412150A100A08626173657463726F12043130303010904E1A40E812FBA1D50115CDDD534C7675B838E6CE7169CDD5AD312182696CABB76F05B82B9557EE1A2842A5579B363F0A5F2E487F7228A81FBF81E125E58F264A29DA07")
    tx, err := decoder.Decode(anyTxBytes)
    // Handle the error and work with tx
    fmt.Pritnln(string(tx.MarshalToJSON()))
}

func DecodeBase64EncodedTx() {
    decoder := cosmostxdecoder.DefaultDecoder
    anyBase64EncodedTx := "CowGCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTkyM3B6MDNtaGphenRnY3YzZ2V5MGhqMGFtd3gwMmR5c2thdTUyCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMXh3ZDNrOHh0ZXJkZWZ0M254cWc5MnN6aHB6NnZ4NDNxc3BkcHc2CpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTh5bGNoZ214eXBodzNjdHNsNzVuNTN1amVxdWttbWFnMm42eDNmCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZzOHI2enhtcjVuYzg2ajhjcGNtam1jY2Y4czJjYWZ4enQ1YWxxCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZqYTVuc3h6N2dzcXc0emNjdXV5OHI3cGpuam1jN2RzY2RsMnZ6EmwKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQL6Inwmwwd0nDUwtu9S8U0E+TU86f92eeo/ZUJfq+O1tRIECgIIARieHRIXChEKCGJhc2V0Y3JvEgUxNTAwMBDwkwkaQDfjeHPhtkdkDG3JyqECSN2DTIZeTC3Z2dK82HL1qshIH6dvMvT2JP4NGhmcQW/JK97sZ+FMdxe98GJxQNLfZfk="
    tx, err := decoder.DecodeBase64(anyBase64EncodedTx)
    // Handle the error and work with tx
    fmt.Pritnln(string(tx.MarshalToJSON()))
}
```

### `NewDecoder`

You can also create your own Decoder and register modules interfaces of your choices to reduce the compiled size:

```go
package yourpackage

import (
    "fmt"

    "github.com/cosmos/cosmos-sdk/codec/types"
    "github.com/cosmos/cosmos-sdk/std"
    banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

    "github.com/calvinlauco/cosmostxdecoder"
)

func DecodeTxBytes() {
    decoder := cosmostxdecoder.NewDecoder()
    // Register only the interfaces of your interest.
    decoder.RegisterInterfaces(RegisterInterfaces)

    anyBase64EncodedTx := "CowGCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTkyM3B6MDNtaGphenRnY3YzZ2V5MGhqMGFtd3gwMmR5c2thdTUyCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMXh3ZDNrOHh0ZXJkZWZ0M254cWc5MnN6aHB6NnZ4NDNxc3BkcHc2CpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTh5bGNoZ214eXBodzNjdHNsNzVuNTN1amVxdWttbWFnMm42eDNmCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZzOHI2enhtcjVuYzg2ajhjcGNtam1jY2Y4czJjYWZ4enQ1YWxxCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZqYTVuc3h6N2dzcXc0emNjdXV5OHI3cGpuam1jN2RzY2RsMnZ6EmwKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQL6Inwmwwd0nDUwtu9S8U0E+TU86f92eeo/ZUJfq+O1tRIECgIIARieHRIXChEKCGJhc2V0Y3JvEgUxNTAwMBDwkwkaQDfjeHPhtkdkDG3JyqECSN2DTIZeTC3Z2dK82HL1qshIH6dvMvT2JP4NGhmcQW/JK97sZ+FMdxe98GJxQNLfZfk="
    tx, err := decoder.DecodeBase64(anyBase64EncodedTx)
    // Handle the error and work with tx
    fmt.Pritnln(string(tx.MarshalToJSON()))
}

func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
}
```

### Work with decoded transaction

To facilitate transaction parsing, there is a transaction struct [CosmosTx](./tx.go).

```go
package yourpackage

import (
    "fmt"
    "encoding/json"

    "github.com/cosmos/cosmos-sdk/codec/types"
    "github.com/cosmos/cosmos-sdk/std"
    banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

    "github.com/calvinlauco/cosmostxdecoder"
)

func DecodeTxBytes() {
    decoder := cosmostxdecoder.NewDecoder()
    // Register only the interfaces of your interest.
    decoder.RegisterInterfaces(RegisterInterfaces)

    anyBase64EncodedTx := "CowGCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTkyM3B6MDNtaGphenRnY3YzZ2V5MGhqMGFtd3gwMmR5c2thdTUyCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMXh3ZDNrOHh0ZXJkZWZ0M254cWc5MnN6aHB6NnZ4NDNxc3BkcHc2CpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTh5bGNoZ214eXBodzNjdHNsNzVuNTN1amVxdWttbWFnMm42eDNmCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZzOHI2enhtcjVuYzg2ajhjcGNtam1jY2Y4czJjYWZ4enQ1YWxxCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZqYTVuc3h6N2dzcXc0emNjdXV5OHI3cGpuam1jN2RzY2RsMnZ6EmwKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQL6Inwmwwd0nDUwtu9S8U0E+TU86f92eeo/ZUJfq+O1tRIECgIIARieHRIXChEKCGJhc2V0Y3JvEgUxNTAwMBDwkwkaQDfjeHPhtkdkDG3JyqECSN2DTIZeTC3Z2dK82HL1qshIH6dvMvT2JP4NGhmcQW/JK97sZ+FMdxe98GJxQNLfZfk="
    tx, err := decoder.DecodeBase64(anyBase64EncodedTx)

    jsonBytes, err := tx.MarshalToJSON()
    var cosmosTx *cosmostxdecoder.CosmosTx
    err := json.Unmarshal(jsonBytes, &cosmosTx)
    // Handle the error and work with CosmosTx struct
   fmt.Pritnln(cosmosTx) 
}

func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
```

## Cosmos Stargate Compatibility

This package is based on Cosmos SDK `v0.40.0`. Release versions are suffixed with Cosmos SDK rc version for compatibility.

For example, if the transaction comes from a chain using Cosmos SDK `v0.40.0-rc1`:
```bash
go checkout release/v0.0.1-rc1
```
```
# go.mod

...
require (
    github.com/calvinlauco/cosmostxdecode v0.0.1-rc3
    ...
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

# Replace `rc1` of the following two lines with your desired Cosmos SDK rc version

replace github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.40.0-rc1

replace github.com/calvinlauco/cosmostxdecoder => github.com/calvinlauco/cosmostxdecoder v0.0.1-rc1
```

## License

This library is published under [MIT license](./LICENSE).
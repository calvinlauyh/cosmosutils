package main

import (
	"encoding/base64"
	"fmt"
	"os"

	. "github.com/calvinlauco/cosmostxdecoder"
	"github.com/thatisuday/commando"
)

func main() {

	// configure commando
	commando.
		SetExecutableName("decode-cosmospubkey").
		SetVersion("1.0.0").
		SetDescription("This tool convert cosmos consensus pubkey/ tendermint pubkey to cosmos consensus address")

	// configure info command
	commando.
		Register(nil).
		SetShortDescription("Convert cosmos consensus pubkey/ tendermint pubkey to cosmos consensus address").
		SetDescription("This command convert cosmos consensus pubkey/ tendermint pubkey to cosmos consensus address with cosmos bech prefix provided").
		AddArgument("pubkey", "pubkey for cosmos consensus pubkey; add '-t' for tendermint pubkey", "").
		AddFlag("prefix,p", "bech32 prefix for Consensus Address", commando.String, "tcrocnclcons"). // required
		AddFlag("tmpubkey,t", "tendermint pubkey for pubkey format", commando.Bool, false).          // required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			var err error
			tmpubkey_bool, _ := flags["tmpubkey"].Value.(bool)
			prefix, _ := flags["prefix"].Value.(string)
			pubkey := args["pubkey"].Value
			var consensusNodeAddress string
			if !tmpubkey_bool {
				consensusNodeAddress, err = ConsensusNodeAddressFromPubKey(prefix, pubkey)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error: Invalid pubkey or prefix")
					os.Exit(1)
				}
			} else {
				tmpubkey, _ := base64.StdEncoding.DecodeString(pubkey)
				consensusNodeAddress, err = ConsensusNodeAddressFromTmPubKey(prefix, tmpubkey)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error: Invalid tm pubkey or prefix")
					os.Exit(1)
				}
			}
			fmt.Printf("Bech32 Validator Consensus Address: %s\n", consensusNodeAddress)
		})

	// parse command-line arguments
	commando.Parse(nil)

}

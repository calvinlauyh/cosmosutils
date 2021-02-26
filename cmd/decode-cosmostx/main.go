package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	. "github.com/calvinlauyh/cosmosutils"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Error: Missing transaction hex")
		printUsage(os.Stderr)
		os.Exit(1)
	}

	txHex := os.Args[1]
	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Invalid transaction hex")
		printUsage(os.Stderr)
		os.Exit(1)
	}

	decoder := DefaultDecoder
	tx, err := decoder.Decode(txBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Invalid transaction")
		printUsage(os.Stderr)
		os.Exit(1)
	}

	jsonBytes, err := tx.MarshalToJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Cannot encode transaction to JSON")
		printUsage(os.Stderr)
		os.Exit(1)
	}
	fmt.Println(string(jsonBytes))
	os.Exit(0)
}

func printUsage(w io.Writer) {
	fmt.Fprintln(w, `Usage:
decode-cosmostx [transaction hex]`)
}

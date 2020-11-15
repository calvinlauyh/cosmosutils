package main

import (
	"fmt"
	"io"
	"os"

	. "github.com/calvinlauco/cosmostxdecoder"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Error: Missing base64 transaction string")
		printUsage(os.Stderr)
		os.Exit(1)
	}

	base64Tx := os.Args[1]

	decoder := DefaultDecoder
	tx, err := decoder.DecodeBase64(base64Tx)
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
decode-cosmosbase64tx [base64 encoded transaction]`)
}

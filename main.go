package main

import (
	"flag"
	"fmt"
)

// CLI flags
var (
	encryptFlag    = flag.Bool("encrypt", false, "boolean flag to encrypt")
	decryptFlag    = flag.Bool("decrypt", false, "boolean flag to decrypt")
	singleFileName = flag.String("s", "", "encrypt/decrypt a single file. Expects a file name")
)

func main() {
	// Parse cli flags
	flag.Parse()

	// Validate flags
	if !validateCLIFlags() {
		return
	}

}

func validateCLIFlags() bool {
	if *encryptFlag && *decryptFlag {
		fmt.Println("You can't encrypt and decrypt at the same time")
		fmt.Println("Terminating")
		return false
	}
	if (*singleFileName != "") && (!*encryptFlag && !*decryptFlag) {
		fmt.Println("Need to encrypt or decrypt the file")
		fmt.Println("Please add --encrypt or --decrypt")
		fmt.Println("Terminating")
		return false
	}
	return true
}

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/frichetten/GoEncryptTheCloud/cryptography"
	"github.com/frichetten/GoEncryptTheCloud/fileoperations"
	"github.com/frichetten/GoEncryptTheCloud/userinput"
)

// CLI flags
var (
	encryptFlag    = flag.Bool("encrypt", false, "boolean flag to encrypt")
	decryptFlag    = flag.Bool("decrypt", false, "boolean flag to decrypt")
	singleFileName = flag.String("s", "", "Encrypt/decrypt a single file. Expects a file name")
	outputFileName = flag.String("o", "", "Decrypt a single file and give it this name")
)

func main() {
	// Parse cli flags
	flag.Parse()

	// Validate flags
	if !validateCLIFlags() {
		return
	}

	// Act on the input
	if *singleFileName != "" && *encryptFlag {
		// Encrypt a single file
		if !fileoperations.IsValidFile(*singleFileName) {
			fmt.Println("Invalid file name")
			fmt.Println("Terminating")
			return
		}
		userPassword := userinput.GetEncryptionKey(*encryptFlag)
		encryptionKey := cryptography.SHA256Hash(userPassword)

		// Read the file into memory
		data := fileoperations.ReadFile(*singleFileName)

		// Encrypt
		ciphertext, err := cryptography.Encrypt(data, encryptionKey)
		if err != nil {
			fmt.Println(err)
		}

		// Write to file
		fileoperations.WriteFile(*singleFileName+".enc", ciphertext)

	} else if *singleFileName != "" && *decryptFlag {
		//Decrypt a single file
		if !fileoperations.IsValidFile(*singleFileName) {
			fmt.Println("Invalid file name")
			fmt.Println("Terminating")
			return
		}
		userPassword := userinput.GetEncryptionKey(*encryptFlag)
		encryptionKey := cryptography.SHA256Hash(userPassword)

		// Read the file into memory
		data := fileoperations.ReadFile(*singleFileName)

		//Decrypt
		plaintext, err := cryptography.Decrypt(data, encryptionKey)
		if err != nil {
			fmt.Println(err)
		}

		// Write to file
		if *outputFileName != "" {
			fileoperations.WriteFile(*outputFileName, plaintext)
		} else {
			*singleFileName = strings.TrimSuffix(*singleFileName, ".enc")
			fileoperations.WriteFile(*singleFileName, plaintext)
		}

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
	if *decryptFlag && *outputFileName != "" && *singleFileName == "" {
		fmt.Println("If you are changing the output name I need to know what file to decrypt")
		fmt.Println("Terminating")
		return false
	}
	return true
}

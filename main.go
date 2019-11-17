package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/frichetten/GoEncryptTheCloud/cryptography"
	"github.com/frichetten/GoEncryptTheCloud/fileoperations"
	"github.com/frichetten/GoEncryptTheCloud/s3"
	"github.com/frichetten/GoEncryptTheCloud/userinput"
)

// CLI flags
var (
	encryptFlag     = flag.Bool("encrypt", false, "boolean flag to encrypt")
	decryptFlag     = flag.Bool("decrypt", false, "boolean flag to decrypt")
	cloudFlag       = flag.Bool("cloud", false, "boolean flag to upload to AWS S3")
	alterConfigFlag = flag.Bool("change", false, "boolean flag to configure the config file")
	localFlag       = flag.Bool("local", false, "This flag will work with local directories rather "+
		"than s3")
	configFile = flag.String("config", "", "Normal location is "+
		"~/.config/GoEncryptTheCloud/config, this sets a custom one")
	bucketName     = flag.String("bucket", "", "AWS S3 bucket to read/write from. Ensure you have permission")
	directoryName  = flag.String("dir", "", "Directory of files to encrypt")
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

	// Load Config file
	config := fileoperations.Config{}
	if fileoperations.ValidateConfigFileLocation(*configFile) {
		config = fileoperations.ReadConfigFile(*configFile)
	}

	// Act on the input
	if *alterConfigFlag {
		// Allow user to change their config
		userinput.AlterConfigFile(*configFile, config)
	} else if *singleFileName != "" && *encryptFlag {
		encryptFile(*singleFileName, config)
	} else if *singleFileName != "" && *decryptFlag {
		//Decrypt a single file
		userPassword := userinput.GetEncryptionKey(*encryptFlag)
		encryptionKey := cryptography.SHA256Hash(userPassword)

		data := []byte{}
		if *cloudFlag {
			// If we are getting files from the cloud we need to download them
			data = s3.DownloadS3(*singleFileName, "jam")
		} else {
			// Read the file into memory
			data = fileoperations.ReadFile(*singleFileName)
		}

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

	} else if *directoryName != "" && *encryptFlag {
		// Encrypt a directory
		userPassword := userinput.GetEncryptionKey(*encryptFlag)
		encryptionKey := cryptography.SHA256Hash(userPassword)

		err := filepath.Walk(*directoryName, func(path string, info os.FileInfo, err error) error {
			fd, err := os.Stat(path)
			if err != nil {
				panic(err)
			}
			if fd.IsDir() {
				directory := strings.TrimSuffix(path, "/")
				directory = strings.Replace(directory, "/", ".enc/", -1)
				os.Mkdir(directory+".enc", os.ModePerm)
				fmt.Println(directory + ".enc")
			} else {
				// Read the file into memory
				data := fileoperations.ReadFile(path)

				//Encrypt
				ciphertext, err := cryptography.Encrypt(data, encryptionKey)
				if err != nil {
					fmt.Println(err)
				}

				// Write to file
				path = strings.Replace(path, "/", ".enc/", -1)
				fileoperations.WriteFile(path+".enc", ciphertext)
				fmt.Println(path + ".enc")
			}
			return nil
		})
		if err != nil {
			panic(err)
		}

	} else if *directoryName != "" && *decryptFlag {
		// Decrypt a directory
		userPassword := userinput.GetEncryptionKey(*encryptFlag)
		encryptionKey := cryptography.SHA256Hash(userPassword)

		err := filepath.Walk(*directoryName, func(path string, info os.FileInfo, err error) error {
			fd, err := os.Stat(path)
			if err != nil {
				panic(err)
			}
			if fd.IsDir() {
				directory := strings.TrimSuffix(path, "/")
				directory = strings.Replace(directory, ".enc", "", -1)
				os.Mkdir(directory, os.ModePerm)
				fmt.Println(directory)
			} else {
				// Read the file into memory
				data := fileoperations.ReadFile(path)

				// Decrypt
				plaintext, err := cryptography.Decrypt(data, encryptionKey)
				if err != nil {
					panic(err)
				}

				// Write to file
				path = strings.Replace(path, ".enc", "", -1)
				fileoperations.WriteFile(path, plaintext)
				fmt.Println(path)
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
}

func encryptFile(filename string, config fileoperations.Config) {
	if !fileoperations.IsValidFile(filename) {
		fmt.Println("Invalid file name")
		fmt.Println("Terminating")
		return
	}
	userPassword := userinput.GetEncryptionKey(*encryptFlag)
	encryptionKey := cryptography.SHA256Hash(userPassword)

	// Read the file into memory
	data := fileoperations.ReadFile(filename)

	// Encrypt
	ciphertext, err := cryptography.Encrypt(data, encryptionKey)
	if err != nil {
		fmt.Println(err)
	}

	// Write to S3 or file
	if *localFlag {
		fileoperations.WriteFile(*singleFileName+".enc", ciphertext)
	} else {
		s3.WriteS3(config.Bucketname, *singleFileName+".enc", ciphertext)
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
	if *singleFileName != "" && *directoryName != "" {
		fmt.Println("You can't encrypt/decrypt a directory and a single file at the same time")
		fmt.Println("Terminating")
		return false
	}
	return true
}

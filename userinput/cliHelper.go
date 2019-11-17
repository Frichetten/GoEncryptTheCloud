package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/frichetten/GoEncryptTheCloud/fileoperations"
)

// GetEncryptionKey will ask the user for the encryption key
func GetEncryptionKey(isEncrypt bool) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the encryption password: ")
	first, _ := reader.ReadString('\n')
	if !isEncrypt {
		return strings.TrimSuffix(first, "\n")
	}

	fmt.Print("Confirm Password: ")
	second, _ := reader.ReadString('\n')
	if first != second {
		fmt.Println("Encryption keys don't match")
		fmt.Println("Terminating")
		os.Exit(3)
	}
	return strings.TrimSuffix(first, "\n")
}

// AlterConfigFile asks the user what they'd like to put in their config file
func AlterConfigFile(manualConfigFileLocationString string, existingConfig fileoperations.Config) {
	newConfig := fileoperations.Config{}
	fmt.Println("We will walk you through the config process")
	fmt.Println("Leaving option blank will leave the current value")

	newConfig.Bucketname = getConfigOption(
		"bucket name [ "+existingConfig.Bucketname+" ]: ",
		existingConfig.Bucketname,
	)
	newConfig.Region = getConfigOption(
		"region [ "+existingConfig.Region+" ]: ",
		existingConfig.Region,
	)

	fileoperations.UpdateConfigFile(manualConfigFileLocationString, newConfig)
}

func getConfigOption(query string, existingValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(query)
	value, _ := reader.ReadString('\n')
	value = strings.TrimSuffix(value, "\n")
	if value == "" {
		return existingValue
	}
	return value
}

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
func AlterConfigFile(manualConfigFileLocationString string) {
	existingConfig := fileoperations.ReadConfigFile(manualConfigFileLocationString)
	newConfig := fileoperations.Config{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("We will walk you through the config process")
	fmt.Println("Leaving option blank will leave the current value")
	fmt.Print("bucket name [ ", existingConfig.Bucketname, " ]: ")
	bucketname, _ := reader.ReadString('\n')
	bucketname = strings.TrimSuffix(bucketname, "\n")
	if bucketname == "" {
		newConfig.Bucketname = existingConfig.Bucketname
	} else {
		newConfig.Bucketname = bucketname
	}
	fmt.Print("region [ ", existingConfig.Region, " ]: ")
	region, _ := reader.ReadString('\n')
	region = strings.TrimSuffix(region, "\n")
	if region == "" {
		newConfig.Region = existingConfig.Region
	} else {
		newConfig.Region = region
	}
	fileoperations.UpdateConfigFile(newConfig)
}

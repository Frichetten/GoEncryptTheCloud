package fileoperations

import (
	"io/ioutil"
	"os"
	"strings"
)

// Config is the data structure to hold the config
type Config struct {
	Bucketname string
	Region     string
}

var userHomeDir, _ = os.UserHomeDir()
var configFileLocation = userHomeDir + "/.config/GoEncryptTheCloud/config"

// ValidateConfigFileLocation returns whether a config file exits of not
func ValidateConfigFileLocation(path string) bool {
	if path != "" {
		return IsValidFile(path)
	}
	return IsValidFile(configFileLocation)
}

// CreateConfigFile creates a config file in the regular place
func CreateConfigFile() {
	os.Create(configFileLocation)
}

// ReadConfigFile reads in the contents of the config file
func ReadConfigFile(manualConfigFileLocation string) Config {
	location := ""
	if manualConfigFileLocation != "" {
		location = manualConfigFileLocation
	} else {
		location = configFileLocation
	}

	data, err := ioutil.ReadFile(location)
	if err != nil {
		panic(err)
	}
	return parseConfigFile(data)
}

func parseConfigFile(data []byte) Config {
	// convert to string
	values := strings.Split(string(data), "\n")
	config := Config{}
	for _, item := range values {
		if strings.Index(item, "bucketname") == 0 {
			config.Bucketname = strings.TrimSuffix(item[strings.Index(item, " ")+1:], "\n")
		} else if strings.Index(item, "region") == 0 {
			config.Region = strings.TrimSuffix(item[strings.Index(item, " ")+1:], "\n")
		}
	}

	return config
}

// UpdateConfigFile updates the config file
func UpdateConfigFile(manualConfigFileLocationString string, newConfig Config) {
	data := "bucketname " + newConfig.Bucketname + "\n" +
		"region " + newConfig.Region + "\n"
	if manualConfigFileLocationString != "" {
		WriteFile(manualConfigFileLocationString, []byte(data))
	} else {
		WriteFile(configFileLocation, []byte(data))
	}
}

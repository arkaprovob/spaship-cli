package cmd

import (
	"log"
	"os"
	"path/filepath"
)

func computeConfigDirectory() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.FromSlash(dirname + "/.spaship")
}

func createConfigFile(configDir string) string {
	os.MkdirAll(configDir, os.ModePerm)
	configFilePath := filepath.FromSlash(configDir + "/config")
	var _, err = os.Stat(configFilePath)
	if os.IsNotExist(err) {
		var file, err = os.Create(configFilePath)
		if os.IsExist(err) {
			log.Fatal(err)
			return ""
		}
		defer file.Close()
	}
	return configFilePath
}

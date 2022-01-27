package cmd

import (
	"io"
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

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
func IsDirexists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

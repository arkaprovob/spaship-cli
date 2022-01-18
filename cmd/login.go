/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "use this command for login ",
	Long:  `Login command for spaship cli to login into a server and perform operations`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("login called")
		serverUrl, _ := cmd.Flags().GetString("server")
		accessToken, _ := cmd.Flags().GetString("token")

		if checkArguments(serverUrl, accessToken) == false {
			return
		}

		log.Println("server url is  '" + serverUrl + "' and access token is '" + accessToken + "'")

		bearer := "Bearer " + accessToken

		responseString := authenticate(serverUrl, bearer)

		configDir := computeConfigDirectory()
		configFile := createConfigFile(configDir)
		log.Println("local user directory is '" + configDir + "'")
		log.Println("config file location '" + configFile + "'")
		storeDetailsInConfig(serverUrl, responseString, configFile)
	},
}

func storeDetailsInConfig(serverAddress string, response string, configPath string) bool {
	var responseJson Response
	err := json.Unmarshal([]byte(response), &responseJson)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	conf := new(Configuration).Init()
	//Configuration{Server: serverAddress, AccessToken: responseJson.AccessToken}
	conf.Server = serverAddress
	conf.AccessToken = responseJson.AccessToken
	confJson, _ := json.Marshal(conf)
	log.Println("storing config " + string(confJson))

	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var confList ConfigList
	json.Unmarshal(byteValue, &confList)
	var entry bool = false;
	for i := 0; i < len(confList.ConfigList); i++ {
		if confList.ConfigList[i].Server==serverAddress {
			entry = true
			confList.ConfigList[i].AccessToken =
		}

	}

	return true
}

func authenticate(url string, accessToken string) string {
	req, err := http.NewRequest("GET", url+"/api/validate", nil)
	req.Header.Add("Authorization", accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		errors.New("unable to fetch reponse")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		errors.New("unable to parse response body")
	}
	return string(body)
}

func checkArguments(server string, token string) bool {
	if len(server) == 0 {
		fmt.Println("Please provide server url ")
		return false
	}
	if len(token) == 0 {
		fmt.Println("Please provide token ")
		return false
	}
	return true
}

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

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().String("server", "", "api url")
	loginCmd.PersistentFlags().String("token", "", "token to access the api")

}
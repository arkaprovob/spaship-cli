/*
Copyright © 2022 Arkaprovo Bhattacharjee <apb@live.in>

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
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
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

		if !checkArguments(serverUrl, accessToken) {
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

	//transform response string into Response structure
	var responseJson Response
	err := json.Unmarshal([]byte(response), &responseJson)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	//creates config entry from current login details
	conf := buildConfigObject(serverAddress, responseJson) //confJson, _ := json.Marshal(conf)//log.Println("storing config " + string(confJson))

	//Load config file into memory from local machine
	jsonFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	//transform the content into confList
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var confList ConfigList
	json.Unmarshal(byteValue, &confList)

	//add current config if the file is empty
	if confList.ConfigList == nil {
		log.Println("No entry found!!")
		confList := ConfigList{}
		confList.AddConfig(conf)
	}
	//Cheks for the existing entry match
	entry := checkExistingEntries(&confList, serverAddress, responseJson)

	// when the current config entry not found add it to the list
	if !entry {
		confList.AddConfig(conf)
	}

	// re-write the list into the file
	file, _ := json.MarshalIndent(confList, "", " ")
	_ = ioutil.WriteFile(configPath, file, 0644)

	return true
}

func checkExistingEntries(confList *ConfigList, serverAddress string, responseJson Response) bool {
	var entry bool
	for i := 0; i < len(confList.ConfigList); i++ {
		if confList.ConfigList[i].Server == serverAddress {
			confList.ConfigList[i].Active = true
			confList.ConfigList[i].AccessToken = responseJson.AccessToken
			confList.ConfigList[i].Alias = responseJson.Identifier
			entry = true
		} else {
			confList.ConfigList[i].Active = false
		}

	}
	return entry
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

func buildConfigObject(serverAddress string, responseJson Response) Configuration {
	conf := new(Configuration).Init()
	conf.Server = serverAddress
	conf.AccessToken = responseJson.AccessToken
	conf.Alias = responseJson.Identifier
	conf.Active = true
	return conf
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().String("server", "", "api url")
	loginCmd.PersistentFlags().String("token", "", "token to access the api")
}

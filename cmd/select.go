/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "use this command to select a spaship server ",
	Long: `using select command the user can switch between multiple spaship deployments, 
	however make sure that the credentials are upto date`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("select called")

		alias, _ := cmd.Flags().GetString("alias")

		if len(alias) == 0 {
			fmt.Println("Please select an alias ")
			return
		}

		configDir := computeConfigDirectory()
		configFile := createConfigFile(configDir)

		switchServer(alias, configFile)

	},
}

func switchServer(alias string, configPath string) {

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

	if confList.ConfigList == nil {
		log.Println("no configuration entry found!, please login before switching the server")
	}

	entry := modifyConfig(&confList, alias)

	if !entry {
		log.Println("alias not found")
	} else {
		// re-write the list into the file
		file, _ := json.MarshalIndent(confList, "", " ")
		_ = ioutil.WriteFile(configPath, file, 0644)
		log.Println("switched successfully")
	}

}

func modifyConfig(confList *ConfigList, alias string) bool {
	var entry bool
	for i := 0; i < len(confList.ConfigList); i++ {
		if confList.ConfigList[i].Alias == alias {
			confList.ConfigList[i].Active = true
			entry = true
		} else {
			confList.ConfigList[i].Active = false
		}

	}
	return entry
}

func init() {
	rootCmd.AddCommand(selectCmd)
	selectCmd.PersistentFlags().String("alias", "", "select the alias of spaship server")
}

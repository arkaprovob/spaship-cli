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
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init will create a .spaship fle into the spa root directory",
	Long: `this will generate the .spaship file into the selected directory, 
	when no dir  specified it will generate the mapping file into the current directory `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		dir, _ := cmd.Flags().GetString("dir")
		var err error
		if len(dir) == 0 {
			dir, err = os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("directory not specified, this will generate the mapping into the current directory")
		}

		var wpn string
		var spaName string
		var route string

		fmt.Print("Web property version: v1\n")

		fmt.Print("Web property name: ")
		fmt.Scanf("%s", &wpn)
		for len(wpn) < 1 {
			fmt.Print("Web property name*: ")
			fmt.Scanf("%s", &wpn)
		}

		fmt.Print("SPA name: ")
		fmt.Scanf("%s", &spaName)
		for len(spaName) < 1 {
			fmt.Print("SPA name*: ")
			fmt.Scanf("%s", &spaName)
		}

		fmt.Print("Route: ")
		fmt.Scanf("%s", &route)
		for len(route) < 1 {
			fmt.Print("Route*: ")
			fmt.Scanf("%s", &route)
		}

		var spaShipMapping SpashipMapping = new(SpashipMapping).Init()
		spaShipMapping.WebsiteName = wpn
		spaShipMapping.SpaName = spaName
		spaShipMapping.Route = route

		var continute string = "y"
		var name string
		var updateRestriction string
		var exclude string

		for strings.Contains(continute, "y") {

			fmt.Print("Environment name: ")
			fmt.Scanf("%s", &name)
			for len(name) < 1 {
				fmt.Print("Environment name*: ")
				fmt.Scanf("%s", &name)
			}

			fmt.Print("Update restriction?: ")
			fmt.Scanf("%s", &updateRestriction)
			for len(updateRestriction) < 1 {
				fmt.Print("Update restriction?*: ")
				fmt.Scanf("%s", &updateRestriction)
			}

			fmt.Print("Exclude from environment?: ")
			fmt.Scanf("%s", &exclude)
			for len(exclude) < 1 {
				fmt.Print("Exclude from environment?*: ")
				fmt.Scanf("%s", &exclude)
			}

			fmt.Print("continue? (type y or press enter to continue or n to end): ")
			fmt.Scanf("%s", &continute)

			var env = Environment{name, strings.Contains(updateRestriction, "y"), strings.Contains(exclude, "y")}
			spaShipMapping.AddEnvironment(env)

		}
		file, _ := json.MarshalIndent(spaShipMapping, "", " ")
		mappingPath := filepath.FromSlash(dir + "/.spaship")
		_ = ioutil.WriteFile(mappingPath, file, 0644)
		fmt.Println("fully qualified mapping path is ", mappingPath)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().String("dir", "", "select directory to store config file")

}

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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// packCmd represents the pack command
var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "pack command will pack the distribution and the mapping in a zip file",
	Long:  `use this command to pack the distribution and .spaship mapping together in a zip file and store it in a directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pack called")
		dist, _ := cmd.Flags().GetString("dist")
		mapping, _ := cmd.Flags().GetString("mapping")
		fmt.Print(mapping)
		if len(dist) == 0 {

			fmt.Print("location spa distribution folder is missing, searching in the current directory\n")

			// list the files and dirs within the current directory
			files, err := ioutil.ReadDir(".")
			if err != nil {
				log.Fatalln(err)
			}
			var dirExists bool
			//checks for a pre defined distribution or a build directory
			for _, f := range files {

				if strings.Contains(f.Name(), "dist") || strings.Contains(f.Name(), "build") {

					// get the current directory location
					dir, err := os.Getwd()
					if err != nil {
						log.Fatalln(err)
					}
					dist = filepath.FromSlash(dir + "/" + f.Name())
					fmt.Println("distribution directory found! " + dist)

					distCOntent, _ := IsDirEmpty(dist)
					if distCOntent {
						fmt.Println("distribution directory is empty, exiting from the pack operation")
						return
					}

					dirExists = true
					break
				}

			}

			if !dirExists {
				fmt.Println(`predefined dist directory not found in current directory..
				 provide the distribution directory location manually or 
				 execute the command from correct directory`)
				return

			}

		}

		distExists, _ := IsDirexists(dist)
		distContent, _ := IsDirEmpty(dist)
		if !distExists || distContent {
			fmt.Println("please select the right directory and execute the pack command")
		}

		if len(mapping) == 0 {

			fmt.Println("searching mapping file in the current directory")
			files, err := ioutil.ReadDir(".")
			if err != nil {
				log.Fatalln(err)
			}
			var mappingExists bool

			for _, f := range files {

				if strings.Contains(f.Name(), ".spaship") {

					// get the current directory location
					dir, err := os.Getwd()
					if err != nil {
						log.Fatalln(err)
					}
					mapping = filepath.FromSlash(dir + "/" + f.Name())
					fmt.Println("distribution directory found! " + mapping)
					mappingExists = true
					break
				}

			}
			if !mappingExists {
				fmt.Println("mapping file not found in the current directory, please execute the pack command with the correct location")
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.PersistentFlags().String("dist", "", "select distribution directory")
	packCmd.PersistentFlags().String("mapping", "", "select mapping file location")

}

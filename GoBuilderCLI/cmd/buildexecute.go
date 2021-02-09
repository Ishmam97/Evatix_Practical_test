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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// buildexecuteCmd represents the buildexecute command
var (
	builddir        bool
	copydir         bool
	exe				bool
	source          string
	destination     string
	buildexecuteCmd = &cobra.Command{
		Use:   "buildexecute",
		Short: "build execute",
		Long:  `USE buildexecute [flags] [sourcefilename] [destinationDir] [fileToBuild]`,
		Run: func(cmd *cobra.Command, args []string) {
			if copydir != true {
				fmt.Println("Must use -c [path/to/file]")
				os.Exit(1)

			} else if builddir != true {
				source = args[0]
				_ = copyFile(source, source)
			} else if exe!=true{
				source = args[0]
				destination = args[1]
				_ = copyFile(source, destination)
			}else{
				source = args[0]
				destination = args[1]
				_ = copyFile(source, destination)
				buildCmd:= fmt.Sprintf("go build %s" ,args[2])
				cmd := exec.Command(buildCmd)

				err := cmd.Run()

				if err != nil {
					log.Fatal(err)
				}
			}

		},
	}
)

func init() {
	rootCmd.AddCommand(buildexecuteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildexecuteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildexecuteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buildexecuteCmd.Flags().BoolVarP(&exe, "exe", "e", false, "-e [fileToBuild]")
	buildexecuteCmd.Flags().BoolVarP(&copydir, "copydir", "c", false, "-c [source]")
	buildexecuteCmd.Flags().BoolVarP(&builddir, "builddir", "b", false, "-b [destination]")

}
func copyFile(s string, d string) error {
	input, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	filenameparse := strings.Split(d, "/")
	filename := filenameparse[len(filenameparse)-1]
	filename = strings.Split(filename, ".")[0] + "-copy." + strings.Split(filename, ".")[1]
	filenameparse[len(filenameparse)-1] = filename
	err = ioutil.WriteFile(strings.Join(filenameparse, "/"), input, 0666)
	if err != nil {
		fmt.Println("Error creating", d)
		fmt.Println(err)
		panic(err)
	}
	return err
}

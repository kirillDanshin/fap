// Copyright © 2016 Kirill Danshin <k@guava.by>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "fap build your application",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("quicktemplate").Value.String() == "true" {
			qtcBuild := exec.Command("qtc", "./...")
			if o, err := qtcBuild.CombinedOutput(); err != nil {
				log.Printf("qtc build error: %s\n\n%s", o, err)
			} else {
				log.Printf("qtc build done")
			}
		}
		var goInstall *exec.Cmd
		if cmd.Flag("debug").Value.String() == "true" {
			log.Println("Using debug mode")
			goInstall = exec.Command("go", "install", `-tags="debug"`)
		} else {
			goInstall = exec.Command("go", "install")
		}
		if o, err := goInstall.CombinedOutput(); err != nil {
			log.Printf("go install error: %s\n\n%s", o, err)
		} else {
			log.Printf("go install done")
		}
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	buildCmd.PersistentFlags().BoolP("quicktemplate", "q", true, "Enable qtc compiling")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	buildCmd.Flags().BoolP("debug", "d", false, "Enable debug mode")

}

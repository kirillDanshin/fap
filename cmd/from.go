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
	"os"
	"os/exec"
	"path"

	"github.com/imdario/mergo"
	"github.com/kirillDanshin/fap/lib/templates/initreadme"
	"github.com/kirillDanshin/fap/lib/templates/web"
	"github.com/kirillDanshin/myutils"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// fromCmd represents the from command
var fromCmd = &cobra.Command{
	Use:   "from",
	Short: "Generate from selected template",
	Run: func(cmd *cobra.Command, args []string) {
		packageStruct, _ := web.Generate("main")
		fs := afero.Afero{Fs: afero.NewOsFs()}
		if cmd.Flag("git").Value.String() == "true" {
			exec.Command("git", "init").Run()
			wd, err := os.Getwd()
			myutils.LogFatalError(err)
			readme, err := initreadme.Generate(path.Base(wd), "fasthttprouter")
			myutils.LogFatalError(err)
			mergo.Merge(&packageStruct, readme)
		}
		for filePath, content := range packageStruct {
			exists, _ := fs.Exists(filePath)
			if exists && cmd.Flag("overwrite").Value.String() == "true" || !exists {
				log.Println("Overwriting", filePath)
				dir, _ := path.Split(filePath)
				if len(dir) != 0 {
					os.MkdirAll(dir, 0755)
				}
				fs.WriteFile(filePath, []byte(content), 0755)
			}
		}
		if cmd.Flag("git").Value.String() == "true" {
			exec.Command("git", "add", ".").Run()
			exec.Command("git", "commit", "-m", "init with fap").Run()
		}
	},
}

func init() {
	initCmd.AddCommand(fromCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fromCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	fromCmd.Flags().BoolP("git", "g", true, "init git repo")
	fromCmd.Flags().BoolP("overwrite", "o", false, "overwrite files")

}

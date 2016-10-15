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

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Adds all changes to git, commits it and pushes to origin/master",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Printf("args: %#+v\n", args)
			log.Println("Enter the release message")
		}
		if o, err := exec.Command("git", "add", ".").CombinedOutput(); err != nil {
			log.Printf("git add . error: %s", o)
		}
		if o, err := exec.Command("git", "commit", "-m", args[0]).CombinedOutput(); err != nil {
			log.Printf("git commit error: %s", o)
		}
		if o, err := exec.Command("git", "push", "origin", "master").CombinedOutput(); err != nil {
			log.Printf("git push origin master error: %s", o)
		}
	},
}

func init() {
	RootCmd.AddCommand(releaseCmd)
}

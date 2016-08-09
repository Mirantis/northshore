// Copyright 2016 The NorthShore Authors All rights reserved.
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
	log "github.com/Sirupsen/logrus"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/spf13/cobra"
)

var blueprintFile string
var blueprintCmd = &cobra.Command{
	Use:   "blueprint",
	Short: "Run execution of blueprint",
	Long:  `This command read, parse and process blueprint.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infoln("Blueprint was runned.")
		log.Infoln("PATH -> ", blueprintFile)
		bp, err := blueprint.ParseFile(blueprintFile)
		if err != nil {
			log.Errorln("Parsing error: ", err)
		}
		log.Debugln("BLUEPRINT -> ", bp)
		log.Infoln("Running............")
		blueprint.RunBlueprint(bp)
	},
}

func init() {
	blueprintCmd.Flags().StringVarP(&blueprintFile, "file", "f", "", "Path to blueprint yaml")
	runCmd.AddCommand(blueprintCmd)
}

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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var path string
var blueprintCmd = &cobra.Command{
	Use:   "blueprint",
	Short: "Run execution of blueprint",
	Long:  `This command read, parse and process blueprint.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Blueprint was runned.")
		fmt.Printf("PATH -> %s \n", path)
		pipeline, err := ParseBlueprint(path)
		if err != nil {
			fmt.Printf("Parsing error: %s \n", err)
		}
		fmt.Printf("PIPELINE -> %+v \n", pipeline)
	},
}

type Stage struct {
	Image       string
	Description string
	Ports       []map[string]int
	Variables   map[string]string
}

type Pipeline struct {
	Version     string
	Type        string
	Name        string
	Provisioner string
	Stages      map[string]Stage
}

func ParseBlueprint(path string) (pipeline Pipeline, err error) {
	viper.SetConfigName("pipeline")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	if err != nil {
		return pipeline, fmt.Errorf("Config not found. %s \n", err)
	}

	err = viper.Unmarshal(&pipeline)
	if err != nil {
		return pipeline, fmt.Errorf("Unable to decode into struct, %v", err)
	}
	return pipeline, nil
}

func init() {
	blueprintCmd.Flags().StringVarP(&path, "file", "f", ".", "Path to blueprint yaml")
	runCmd.AddCommand(blueprintCmd)
}

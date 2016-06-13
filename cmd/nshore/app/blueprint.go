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

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

var path string
var blueprintCmd = &cobra.Command{
	Use:   "blueprint",
	Short: "Run execution of blueprint",
	Long:  `This command read, parse and process blueprint.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Blueprint was runned.")
		fmt.Printf("PATH -> %s \n", path)
		bp, err := ParseBlueprint(path)
		if err != nil {
			fmt.Printf("Parsing error: %s \n", err)
			return
		}
		fmt.Printf("BLUEPRINT -> %+v \n", bp)
		fmt.Println("Running...............")
		RunBlueprint(bp)
	},
}

// Stage represents a Blueprint Stage
type Stage struct {
	//Docker image for bootstrap stage
	Image       string `json:"image"`
	Description string `json:"description"`
	//Ports for exposing to host
	Ports []map[string]string `json:"ports"`
	//Environment variables
	Variables map[string]string `json:"variables"`
}

// Blueprint represents a Blueprint
type Blueprint struct {
	//API version for processing blueprint
	Version string `json:"version"`
	//Type of blueprint (pipeline/application)
	Type string `json:"type"`
	Name string `json:"name"`
	//Provisioner type (docker/...)
	Provisioner string           `json:"provisioner"`
	Stages      map[string]Stage `json:"stages"`
}

// ParseBlueprint parses and validates the incoming data
func ParseBlueprint(path string) (bp Blueprint, err error) {
	viper.SetConfigName("pipeline")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	if err != nil {
		return bp, fmt.Errorf("Config not found. %s \n", err)
	}

	err = viper.Unmarshal(&bp)
	if err != nil {
		return bp, fmt.Errorf("Unable to decode into struct, %v", err)
	}
	return bp, nil
}

func RunBlueprint(bp Blueprint) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	for name, stage := range bp.Stages {
		bindings := make(map[nat.Port][]nat.PortBinding)
		for _, ports := range stage.Ports {
			port, _ := nat.NewPort("tcp", ports["fromPort"])
			bindings[port] = []nat.PortBinding{nat.PortBinding{HostIP: "0.0.0.0", HostPort: ports["toPort"]}}
		}

		hostConfig := container.HostConfig{
			PortBindings: bindings,
		}

		config := container.Config{
			Image: bp.Stages[name].Image,
		}
		fmt.Printf("%s -> Config was built. \n", name)

		r, err := cli.ContainerCreate(context.Background(), &config, &hostConfig, nil, name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s -> Container was created. \n", name)

		err = cli.ContainerStart(
			context.Background(),
			r.ID,
			types.ContainerStartOptions{})
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s -> Container was started. \n", name)
		fmt.Printf("%s -> Container ID  %s \n", name, r.ID)
		fmt.Printf("%s -> Warnings: %s \n", name, r.Warnings)
		fmt.Println("======= \t ======= \t ======= \t ======= \t ======= \t ======= \t =======")
	}
}

func init() {
	blueprintCmd.Flags().StringVarP(&path, "file", "f", ".", "Path to blueprint yaml")
	runCmd.AddCommand(blueprintCmd)
}

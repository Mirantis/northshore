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
	"log"
	"strings"

	"github.com/boltdb/bolt"
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
		log.Println("Blueprint was runned.")
		log.Printf("PATH -> %s", path)
		bp, err := ParseBlueprint(path)
		if err != nil {
			log.Fatalf("Parsing error: %s \n", err)
		}
		log.Printf("BLUEPRINT -> %+v", bp)
		log.Println("Running............")
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
		return bp, fmt.Errorf("Config not found. %s", err)
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
	ids := []string{}

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
		log.Printf("%s -> Config was built.", name)

		r, err := cli.ContainerCreate(context.Background(), &config, &hostConfig, nil, name)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("%s -> Container was created.", name)
		ids = append(ids, r.ID)

		err = cli.ContainerStart(
			context.Background(),
			r.ID,
			types.ContainerStartOptions{})
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("%s -> Container was started.", name)
		log.Printf("%s -> Container ID  %s", name, r.ID)
		log.Printf("%s -> Warnings: %s", name, r.Warnings)
	}
	if len(ids) > 0 {
		updateIDs(strings.Join(ids[:], ","))
	}
}

//Update list of containers in DB
//TODO add ability to add one container
func updateIDs(ids string) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bname := []byte("Northshore")
	key := []byte("containers")
	value := []byte(ids)
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bname)
		if err != nil {
			return fmt.Errorf("Create bucket: %s", err)
		}
		log.Printf("Bucket \"%s\" created\n", bname)
		err = b.Put(key, value)
		if err != nil {
			return err
		}
		log.Printf("Info puted with key \"%s\"\n", key)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	blueprintCmd.Flags().StringVarP(&path, "file", "f", ".", "Path to blueprint yaml")
	runCmd.AddCommand(blueprintCmd)
}

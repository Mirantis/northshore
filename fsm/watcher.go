// Copyright 2016 The NorthShore Authors All rights reserved.
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

package fsm

import (
	"log"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/docker/engine-api/client"
	"golang.org/x/net/context"
)

func Watch(period int, states chan map[string]string) {
	log.Println("Watcher was started...")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	for {
		ids := getIds("my.db")
		if len(ids) == 0 {
			log.Println("No containers for watching.")
			time.Sleep(time.Duration(period) * time.Second)
			continue
		}
		for _, id := range ids {
			res, err := cli.ContainerInspect(context.Background(), id)
			if err != nil {
				log.Println(err)
				continue
			}
			states <- map[string]string{res.Name[1:]: res.State.Status}
			log.Printf(`Container "%s" with id "%s" is in status "%s"`, res.Name, id, res.State.Status)
		}
		time.Sleep(time.Duration(period) * time.Second)
	}
}

func getIds(path string) []string {
	containers := []string{}
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	bname := []byte("Northshore")
	key := []byte("containers")

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bname)
		if bucket == nil {
			log.Printf("Bucket %s not found", bname)
			return nil
		}

		k := bucket.Get(key)
		str := string(k[:])
		containers = strings.Split(str, ",")
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return containers
}

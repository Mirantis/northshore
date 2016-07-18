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

	"github.com/Mirantis/northshore/store"
	"github.com/docker/engine-api/client"
	"golang.org/x/net/context"
)

const (
	// DBBucketWatcher defines boltdb bucket for Watcher
	DBBucketWatcher = "Northshore"

	// DBKeyWatcher defines boltdb key for Watcher
	DBKeyWatcher = "containers"
)

// Watch keeps watch over containers
func Watch(period int) {
	log.Println("Watcher was started...")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	for {
		ids := getIds()
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
			log.Printf(`Container "%s" with id "%s" is in status "%s"`, res.Name, id, res.State.Status)
		}
		time.Sleep(time.Duration(period) * time.Second)
	}
}

func getIds() []string {
	var buf string
	if err := store.Load([]byte(DBBucketWatcher), []byte(DBKeyWatcher), buf); err != nil {
		//Workaround for run local without containers in DB
		log.Print(err)
		return make([]string, 0)
	}

	return strings.Split(buf, ",")
}

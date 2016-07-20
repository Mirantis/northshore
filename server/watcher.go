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

package server

import (
	"strings"
	"time"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/Mirantis/northshore/store"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"golang.org/x/net/context"
)

// Watch keeps watch over containers
func Watch(period int) {
	log.Infoln("Watcher was started...")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	for {
		ids := getIds()
		if len(ids) == 0 {
			log.Infoln("No containers for watching.")
			time.Sleep(time.Duration(period) * time.Second)
			continue
		}
		for _, id := range ids {
			res, err := cli.ContainerInspect(context.Background(), id)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Infof(`Container "%s" with id "%s" is in status "%s"`, res.Name, id, res.State.Status)
		}
		time.Sleep(time.Duration(period) * time.Second)
	}
}

func getIds() []string {
	var data string
	log.Debugln("Get ID's from DB")
	if err := store.Load([]byte(blueprint.DBBucketWatcher), []byte(blueprint.DBKeyWatcher), &data); err != nil {
		//Workaround for run local without containers in DB
		log.Errorln(err)
		return make([]string, 0)
	}
	return strings.Split(data, ",")
}

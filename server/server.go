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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/Mirantis/northshore/fsm"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

var bpl blueprint.BP

func Run(bpPath string) {
	r := mux.NewRouter()

	uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
	uiAPI1.HandleFunc("/", UiAPI1RootHandler).Methods("GET")

	uiAPI1.HandleFunc("/blueprints", blueprints).Methods("GET", "POST")

	ui := r.PathPrefix("/ui").Subrouter().StrictSlash(true)
	ui.PathPrefix("/{uiDir:(app)|(assets)|(node_modules)}").Handler(http.StripPrefix("/ui", http.FileServer(http.Dir("ui/"))))
	ui.HandleFunc("/{s}", UiIndexHandler)
	ui.HandleFunc("/", UiIndexHandler)

	// with 'northshore run local', you can got to http://localhost:8998/ and see a list of
	// what is in static ... if you put index.html in there, it'll be returned.
	// NB: do not put /static in the path, that'll get you a 404.
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bname := []byte("MyBucket")
	key := []byte("answer")
	value := []byte("42")
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

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bname)
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found", bname)
		}

		v := bucket.Get(key)
		log.Printf("Get value by key \"%s\": v => \"%s\" \n", key, v)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	//TODO: draft for demo
	states := make(chan map[string]string, 3)

	bp, err := blueprint.ParseBlueprint(bpPath)
	if err != nil {
		log.Fatalf("Parsing error: %s \n", err)
	}

	var stages []string
	for k := range bp.Stages {
		stages = append(stages, k)
	}

	go func(c chan map[string]string) {
		pl := fsm.NewBlueprintPipeline(stages)
		bpl = blueprint.BP{&bp, pl}
		bpl.Start()

		time.Sleep(time.Second * 9)

		for _, s := range stages {
			time.Sleep(time.Second * 3)
			v := map[string]fsm.StageState{s: fsm.StageStateCreated}
			log.Println("#PL_UPDATE (INITIALIZATION STATE CREATED)!!!", v)
			bpl.Update(v)
		}

		for {
			state := <-c
			log.Printf("CHANNEL IN FSM GOROUTINE -> %v", state)

			for k, v := range state {
				switch v {
				case "running":
					vv := map[string]fsm.StageState{k: fsm.StageStateRunning}
					log.Println("#PL_UPDATE!!!", vv)
					bpl.Update(vv)
				case "exited":
					vv := map[string]fsm.StageState{k: fsm.StageStateStopped}
					log.Println("#PL_UPDATE!!!", vv)
					bpl.Update(vv)
				default:
					log.Println("DEFAULT CASE IN SWITCH!!!")
				}
			}
		}
	}(states)

	//Update frequency for watcher in seconds
	period := 3
	go func(c chan map[string]string) {
		fsm.Watch(period, c)
	}(states)

	log.Println("Listening at port 8998")
	http.ListenAndServe(":8998", r)
}

func UiAPI1RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"version": 1})
}

func UiIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("#uiIndexHandler")
	http.ServeFile(w, r, "ui/index.html")
}

func blueprints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	o := map[string]interface{}{
		"data": []blueprint.BP{
			bpl,
		},
		"meta": map[string]interface{}{
			"info": "blueprints",
		},
	}

	json.NewEncoder(w).Encode(o)
}

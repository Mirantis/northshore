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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var demoBlueprintPath string
var demoBpPl *BlueprintPipeline

// demoBlueprintCmd represents the "demo-blueprint" command
var demoBlueprintCmd = &cobra.Command{
	Use:   "demo-blueprint",
	Short: "Run execution of blueprint",
	Long:  `This command read, parse and process blueprint.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run Blueprint")
		fmt.Printf("PATH -> %s \n", path)
		bp, err := ParseBlueprint(path)
		if err != nil {
			fmt.Printf("Parsing error: %s \n", err)
		}
		fmt.Printf("BLUEPRINT -> %+v \n", bp)
	},
}

// demoFSMCmd represents the "demo-fsm" command
var demoFSMCmd = &cobra.Command{
	Use:   "demo-fsm",
	Short: "Demo FSM",
	Long:  `Run the Blueprint Pipeline thru states`,
	Run: func(cmd *cobra.Command, args []string) {

		stages := []string{"Stage A", "Stage B"}
		pl := NewBlueprintPipeline(stages)

		pl.Start()
		pl.Update(map[string]StageState{"Stage B": StageStateRunning})
		pl.Update(map[string]StageState{"Stage A": StageStateRunning, "Stage B": StageStatePaused})
		pl.Update(map[string]StageState{"Stage B": StageStateRunning})

	},
}

// demoCmd represents the "demo" command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Run demo",
	Long: `Run demo on local server

	The local server binds localhost:8998.
	Demo Blueprint Pipeline goes thru states`,
	Run: func(cmd *cobra.Command, args []string) {

		/* Run Blueprint */
		log.Println("#run_blueprint")
		log.Printf("PATH -> %s \n", path)
		bp, err := ParseBlueprint(path)
		if err != nil {
			log.Printf("Parsing error: %s \n", err)
		}
		log.Printf("BLUEPRINT -> %+v \n", bp)

		/* Run States */
		log.Println("#run_states")
		var stages []string
		for k := range bp.Stages {
			stages = append(stages, k)
		}

		ss := []map[string]StageState{
			{stages[0]: StageStatePaused},
			{stages[0]: StageStateRunning, stages[1]: StageStatePaused},
			{stages[1]: StageStateRunning},
		}

		go func() {
			for {
				demoBpPl = NewBlueprintPipeline(stages)
				demoBpPl.Start()

				for _, s := range stages {
					time.Sleep(time.Second * 1)
					v := map[string]StageState{s: StageStateCreated}
					log.Println("#pl-update", v)
					demoBpPl.Update(v)
				}
				for _, s := range stages {
					time.Sleep(time.Second * 1)
					v := map[string]StageState{s: StageStateRunning}
					log.Println("#pl-update", v)
					demoBpPl.Update(v)
				}
				for _, v := range ss {
					time.Sleep(time.Second * 3)
					log.Println("#pl-update", v)
					demoBpPl.Update(v)
				}
			}
		}()

		/* Run DB */
		/*
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
		*/

		/* Run local server */
		log.Println("#run_local_server")
		r := mux.NewRouter()

		uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
		uiAPI1.HandleFunc("/", uiAPI1RootHandler).Methods("GET")

		uiAPI1.HandleFunc("/action", demouiAPI1ActionHandler).Methods("GET", "POST")
		uiAPI1.HandleFunc("/status", demouiAPI1StatusHandler).Methods("GET")

		ui := r.PathPrefix("/ui").Subrouter().StrictSlash(true)
		ui.PathPrefix("/{uiDir:(app)|(assets)|(node_modules)}").Handler(http.StripPrefix("/ui", http.FileServer(http.Dir("ui/"))))
		ui.HandleFunc("/{s}", uiIndexHandler)
		ui.HandleFunc("/", uiIndexHandler)

		// with 'nshore run local', you can got to http://localhost:8998/ and see a list of
		// what is in static ... if you put index.html in there, it'll be returned.
		// NB: do not put /static in the path, that'll get you a 404.
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

		log.Println("Listening at port 8998")
		http.ListenAndServe(":8998", r)

	},
}

func init() {
	demoBlueprintCmd.Flags().StringVarP(&demoBlueprintPath, "file", "f", ".", "Path to blueprint yaml")
	runCmd.AddCommand(demoBlueprintCmd)
	runCmd.AddCommand(demoFSMCmd)
	runCmd.AddCommand(demoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// localCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// localCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func demouiAPI1ActionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "demouiAPI1ActionHandler"})
}

func demouiAPI1StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(demoBpPl)
}

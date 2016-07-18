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
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/Mirantis/northshore/fsm"
	"github.com/Mirantis/northshore/server"
	"github.com/Mirantis/northshore/store"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var demoBlueprintPath string

// demoBlueprintCmd represents the "demo-blueprint" command
var demoBlueprintCmd = &cobra.Command{
	Use:   "demo-blueprint",
	Short: "Run execution of blueprint",
	Long:  `This command read, parse and process blueprint.`,
	Run: func(cmd *cobra.Command, args []string) {
		bp, err := blueprint.ParseBlueprint(demoBlueprintPath)
		if err != nil {
			log.WithError(err).Fatal("Blueprint parsing error")
		}
		log.WithFields(log.Fields{
			"path":      demoBlueprintPath,
			"blueprint": bp,
		}).Info("Blueprint parsing")
	},
}

// demoFSMCmd represents the "demo-fsm" command
var demoFSMCmd = &cobra.Command{
	Use:   "demo-fsm",
	Short: "Demo FSM",
	Long:  `Run the Blueprint FSM thru states`,
	Run: func(cmd *cobra.Command, args []string) {

		pl := fsm.NewBlueprintFSM(map[string]fsm.StageState{"Stage A": fsm.StageStateNew, "Stage B": fsm.StageStateNew})

		pl.Update(map[string]fsm.StageState{"Stage B": fsm.StageStateRunning})
		pl.Update(map[string]fsm.StageState{"Stage A": fsm.StageStateRunning, "Stage B": fsm.StageStatePaused})
		pl.Update(map[string]fsm.StageState{"Stage B": fsm.StageStateRunning})

	},
}

// demoCmd represents the "demo" command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Run demo",
	Long: `Run demo on local server.

The local server binds localhost:8998.
Demo Blueprint Pipeline goes thru states.`,
	Run: func(cmd *cobra.Command, args []string) {

		/* Run Blueprint */
		bp, err := blueprint.ParseBlueprint(demoBlueprintPath)
		if err != nil {
			log.WithError(err).Fatal("Blueprint parsing error")
		}
		log.WithFields(log.Fields{
			"path":      demoBlueprintPath,
			"blueprint": bp,
		}).Info("Blueprint parsing")

		/* Run States */
		log.Info("Run blueprint states")
		var stages []string
		for k := range bp.Stages {
			stages = append(stages, k)
		}

		ss := []map[string]fsm.StageState{
			{stages[0]: fsm.StageStatePaused},
			{stages[0]: fsm.StageStateRunning, stages[1]: fsm.StageStatePaused},
			{stages[1]: fsm.StageStateRunning},
		}

		go func() {
			for {
				demoBp := blueprint.NewBP(&bp)
				time.Sleep(time.Second * 3)

				for _, s := range stages {
					time.Sleep(time.Second * 3)
					v := map[string]fsm.StageState{s: fsm.StageStateCreated}
					log.Infoln("#pl-update", v)
					demoBp.Update(v)
				}
				for _, s := range stages {
					time.Sleep(time.Second * 3)
					v := map[string]fsm.StageState{s: fsm.StageStateRunning}
					log.Infoln("#pl-update", v)
					demoBp.Update(v)
				}
				for _, v := range ss {
					time.Sleep(time.Second * 3)
					log.Infoln("#pl-update", v)
					demoBp.Update(v)
				}

				store.Delete([]byte(blueprint.DBBucketBlueprints), demoBp.UUID.Bytes())
			}
		}()

		/* Run local server */
		log.Info("Run local server")
		r := mux.NewRouter()

		uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
		uiAPI1.HandleFunc("/", server.UIAPI1RootHandler).Methods("GET")

		uiAPI1.HandleFunc("/action", demouiAPI1ActionHandler).Methods("GET", "POST")
		uiAPI1.HandleFunc("/blueprints", server.UIAPI1BlueprintsHandler).Methods("GET")
		uiAPI1.HandleFunc("/errors", demouiAPI1ErrorsHandler).Methods("GET", "POST")

		ui := r.PathPrefix("/ui").Subrouter().StrictSlash(true)
		ui.PathPrefix("/{uiDir:(app)|(assets)|(dist)|(node_modules)}").Handler(http.StripPrefix("/ui", server.NoDirListing(http.FileServer(http.Dir("ui/")))))
		ui.HandleFunc("/{_:.*}", server.UIIndexHandler)

		r.HandleFunc("/{_:.*}", server.UIIndexHandler)

		httpListen := viper.GetString("HTTPListen")
		log.WithField("httpListen", httpListen).Infoln("#http", "Listen And Serve")
		http.ListenAndServe(httpListen, r)
	},
}

func init() {
	/* Init DB */
	os.Remove("demo.db")
	viper.Set("BoltDBPath", "demo.db")

	/* Init Logger */
	viper.Set("LogLevel", log.DebugLevel)
	log.SetLevel(log.DebugLevel)

	/* Init cobra */
	demoBlueprintCmd.Flags().StringVarP(&demoBlueprintPath, "file", "f", "", "Path to blueprint yaml")
	demoCmd.Flags().StringVarP(&demoBlueprintPath, "file", "f", "", "Path to blueprint yaml")
	runCmd.AddCommand(demoBlueprintCmd)
	runCmd.AddCommand(demoFSMCmd)
	runCmd.AddCommand(demoCmd)
}

func demouiAPI1ActionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	o := map[string]interface{}{
		"data": []map[string]interface{}{
			{"details": "Details 1"},
			{"details": "Details 2"},
		},
		"meta": map[string]interface{}{
			"info": "demouiAPI1ActionHandler",
		},
	}

	json.NewEncoder(w).Encode(o)
}

func demouiAPI1ErrorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(500)

	o := map[string]interface{}{
		"data": []map[string]interface{}{},
		"errors": []map[string]interface{}{
			{
				"details": "Error details 1",
				"title":   "Error title 1",
			},
			{
				"details": "Details of Error 2",
				"meta": map[string]interface{}{
					"info": "meta info of Error 2",
				},
			},
		},
		"meta": map[string]interface{}{
			"info": "demouiAPI1ErrorsHandler",
		},
	}

	json.NewEncoder(w).Encode(o)
}

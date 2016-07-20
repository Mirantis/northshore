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
	"github.com/Mirantis/northshore/server"
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
		overrideSettings()

		bp, err := blueprint.ParseFile(demoBlueprintPath)
		if err != nil {
			log.WithError(err).Fatal("Blueprint parsing error")
		}
		log.WithFields(log.Fields{
			"path":      demoBlueprintPath,
			"blueprint": bp,
		}).Info("Blueprint parsing")
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
		overrideSettings()

		/* Run Blueprint */
		bp, err := blueprint.ParseFile(demoBlueprintPath)
		if err != nil {
			log.WithError(err).Fatal("Blueprint parsing error")
		}
		log.WithFields(log.Fields{
			"path":      demoBlueprintPath,
			"blueprint": bp,
		}).Info("Blueprint parsing")

		bp.Save()
		log.Infoln("#bp_parsed", blueprintStates(bp), bp.State)

		/* Run States */
		log.Info("Run blueprint states")

		ss := [][]blueprint.StageState{
			{blueprint.StageStatePaused},
			{blueprint.StageStateRunning, blueprint.StageStatePaused},
			{blueprint.StageStateRunning, blueprint.StageStateRunning},
			{"badstate", blueprint.StageStateRunning},
		}

		go func() {
			for {

				for i, s := range bp.Stages {
					time.Sleep(time.Second * 3)
					s.State = blueprint.StageStateCreated
					bp.Stages[i] = s
					bp.Save()
					log.Infoln("#bp_update", blueprintStates(bp), bp.State)
				}

				for i, s := range bp.Stages {
					time.Sleep(time.Second * 3)
					s.State = blueprint.StageStateRunning
					bp.Stages[i] = s
					bp.Save()
					log.Infoln("#bp_update", blueprintStates(bp), bp.State)
				}

				for _, v := range ss {
					time.Sleep(time.Second * 3)

					idx := 0
					for i, s := range bp.Stages {
						if idx < len(v) {
							s.State = v[idx]
							bp.Stages[i] = s
						}
						idx++
					}
					bp.Save()
					log.Infoln("#bp_update", blueprintStates(bp), bp.State)
				}

				time.Sleep(time.Second * 3)
				for i, s := range bp.Stages {
					s.State = blueprint.StageStateNew
					bp.Stages[i] = s
				}
				bp.Save()
				log.Infoln("#bp_update", blueprintStates(bp), bp.State)
			}
		}()

		/* Run local server */
		log.Info("Run local server")
		r := mux.NewRouter()

		uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
		uiAPI1.HandleFunc("/", server.UIAPI1RootHandler).Methods("GET")

		uiAPI1.HandleFunc("/action", demouiAPI1ActionHandler).Methods("GET", "POST")
		uiAPI1.HandleFunc("/blueprints", server.UIAPI1BlueprintsHandler).Methods("GET")
		uiAPI1.HandleFunc("/blueprints", server.UIAPI1BlueprintsCreateHandler).Methods("POST")
		uiAPI1.HandleFunc("/blueprints/{id}", server.UIAPI1BlueprintsDeleteHandler).Methods("DELETE")
		uiAPI1.HandleFunc("/blueprints/{id}", server.UIAPI1BlueprintsIDHandler).Methods("GET")
		uiAPI1.HandleFunc("/errors", demouiAPI1ErrorsHandler).Methods("GET", "POST")

		ui := r.PathPrefix("/ui").Subrouter().StrictSlash(true)
		ui.PathPrefix("/{uiDir:(app)|(assets)|(dist)|(node_modules)}").Handler(
			http.StripPrefix("/ui", server.NoDirListing(
				http.FileServer(http.Dir(viper.GetString("UIRoot"))))))
		ui.HandleFunc("/{_:.*}", server.UIIndexHandler)

		r.HandleFunc("/{_:.*}", server.UIIndexHandler)

		ip := viper.GetString("ServerIP")
		port := viper.GetString("ServerPort")
		addr := ip + ":" + port
		log.WithField("address", addr).Infoln("#http", "Listen And Serve")
		log.WithField("UIRoot", viper.GetString("UIRoot")).Infoln("#viper")
		http.ListenAndServe(addr, r)
	},
}

func init() {
	/* Init cobra */
	demoBlueprintCmd.Flags().StringVarP(&demoBlueprintPath, "file", "f", "", "Path to blueprint yaml")
	demoCmd.Flags().StringVarP(&demoBlueprintPath, "file", "f", "", "Path to blueprint yaml")
	demoCmd.Flags().StringVarP(&UIRoot, "ui", "u", "./ui", "Path to UI root directory")
	viper.BindPFlag("UIRoot", demoCmd.Flags().Lookup("ui"))

	runCmd.AddCommand(demoBlueprintCmd)
	runCmd.AddCommand(demoCmd)
}

func demouiAPI1ActionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	ans := map[string]interface{}{
		"data": []map[string]interface{}{
			{"details": "Details 1"},
			{"details": "Details 2"},
		},
		"meta": map[string]interface{}{
			"info": "demouiAPI1ActionHandler",
		},
	}

	log.WithFields(log.Fields{
		"request":  r,
		"response": ans,
	}).Debugln("#http,#demouiAPI1ActionHandler")
	json.NewEncoder(w).Encode(ans)
}

func demouiAPI1ErrorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(500)

	ans := map[string]interface{}{
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

	log.WithFields(log.Fields{
		"request":  r,
		"response": ans,
	}).Debugln("#http,#demouiAPI1ErrorsHandler")
	json.NewEncoder(w).Encode(ans)
}

func blueprintStates(bp blueprint.Blueprint) (ss []blueprint.StageState) {
	for _, s := range bp.Stages {
		ss = append(ss, s.State)
	}
	return ss
}

func overrideSettings() {
	/* Init DB */
	os.Remove("demo.db")
	viper.Set("BoltDBPath", "demo.db")

	/* Init Logger */
	viper.Set("LogLevel", log.DebugLevel)
	log.SetLevel(log.DebugLevel)
}

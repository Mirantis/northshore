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
	"net/http"
	"os"
	"path"
	"time"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/Mirantis/northshore/server"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gingonic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var demoBlueprintFile string

// demoBlueprintCmd represents the "demo-blueprint" command
var demoBlueprintCmd = &cobra.Command{
	Use:   "demo-blueprint",
	Short: "Run execution of blueprint",
	Long:  `This command read, parse and process blueprint.`,
	Run: func(cmd *cobra.Command, args []string) {
		overrideSettings()

		bp, err := blueprint.ParseFile(demoBlueprintFile)
		if err != nil {
			log.WithError(err).Fatal("Blueprint parsing error")
		}
		log.WithFields(log.Fields{
			"path":      demoBlueprintFile,
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
		bp, err := blueprint.ParseFile(demoBlueprintFile)
		if err != nil {
			log.WithError(err).Fatal("Blueprint parsing error")
		}
		log.WithFields(log.Fields{
			"path":      demoBlueprintFile,
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

		r := gin.Default()
		r.NoRoute(func(c *gin.Context) {
			c.File(path.Join(viper.GetString("UIRoot"), "index.html"))
		})

		ui := r.Group("/ui")
		ui.Static("/app", path.Join(viper.GetString("UIRoot"), "/app"))
		ui.Static("/assets", path.Join(viper.GetString("UIRoot"), "/assets"))
		ui.Static("/dist", path.Join(viper.GetString("UIRoot"), "/dist"))
		ui.Static("/node_modules", path.Join(viper.GetString("UIRoot"), "/node_modules"))

		r.GET("/api", server.APIRootHandler)
		r.POST("/ui/api/v1/parse/blueprint", server.APIParseBlueprintHandler)
		api1 := api2go.NewAPIWithRouting(
			"/ui/api/v1",
			api2go.NewStaticResolver("/"),
			gingonic.New(r),
		)
		api1.AddResource(blueprint.Blueprint{}, &server.BlueprintResource{})

		// Custom api route
		// https://github.com/manyminds/api2go/issues/256#issuecomment-234923954
		// it couses to stack overflow
		//
		// api1Handler := api1.Handler().(*httprouter.Router)
		// api1Handler.GET("/a1", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	json.NewEncoder(w).Encode(gin.H{"txt": "/a1"})
		// })
		//
		// api1Router := api1.Router().(*routing.HTTPRouter)
		// it works:
		api1Router := api1.Router()
		api1Router.Handle("GET", "/api/a", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(gin.H{"/api/a": 1})
		})

		addr := fmt.Sprintf(
			"%s:%s",
			viper.GetString("ServerIP"),
			viper.GetString("ServerPort"),
		)
		log.WithFields(log.Fields{
			"addr":   addr,
			"UIRoot": viper.GetString("UIRoot"),
		}).Infoln("#http", "Listen And Serve")
		http.ListenAndServe(addr, r)
	},
}

func init() {
	/* Init cobra */
	demoBlueprintCmd.Flags().StringVarP(&demoBlueprintFile, "file", "f", "", "Path to blueprint yaml")
	demoCmd.Flags().StringVarP(&demoBlueprintFile, "file", "f", "", "Path to blueprint yaml")
	demoCmd.Flags().StringVarP(&UIRoot, "ui", "u", "./ui", "Path to UI root directory")
	viper.BindPFlag("UIRoot", demoCmd.Flags().Lookup("ui"))

	runCmd.AddCommand(demoBlueprintCmd)
	runCmd.AddCommand(demoCmd)
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

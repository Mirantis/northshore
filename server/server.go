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
	"net/http"
	"path"

	"github.com/Mirantis/northshore/blueprint"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gingonic"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/spf13/viper"
)

func Run() {
	go func() {
		Watch(viper.GetInt("WatchPeriod"))
	}()

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.File(path.Join(viper.GetString("UIRoot"), "index.html"))
	})

	ui := r.Group("/ui")
	ui.Static("/app", path.Join(viper.GetString("UIRoot"), "/app"))
	ui.Static("/assets", path.Join(viper.GetString("UIRoot"), "/assets"))
	ui.Static("/dist", path.Join(viper.GetString("UIRoot"), "/dist"))
	ui.Static("/node_modules", path.Join(viper.GetString("UIRoot"), "/node_modules"))

	r.GET("/api", APIRootHandler)
	r.POST("/api/v1/parse/blueprint", APIParseBlueprintHandler)
	api1 := api2go.NewAPIWithRouting(
		"/api/v1",
		api2go.NewStaticResolver("/"),
		gingonic.New(r),
	)
	api1.AddResource(blueprint.Blueprint{}, &BlueprintResource{})

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
}

// APIData writes JSONAPI compatible structure
func APIData(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	json, err := jsonapi.Marshal(v)
	if err != nil {
		APIError(w, err, http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(json))
}

// APIError writes JSONAPI compatible structure
func APIError(w http.ResponseWriter, e error, status int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(gin.H{
		"errors": []gin.H{
			{"detail": e.Error()},
		},
	})
}

// APIRootHandler returns the API description
func APIRootHandler(c *gin.Context) {
	c.JSON(200, gin.H{"/api/v1": gin.H{"version": 1, "type": "jsonapi", "url": "http://jsonapi.org/format/1.0/"}})
}

// APIParseBlueprintHandler creates and stores a blueprint
func APIParseBlueprintHandler(c *gin.Context) {
	ParseBlueprintHandler(c.Writer, c.Request)
}

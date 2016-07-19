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
	"net/http"
	"path"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/Mirantis/northshore/store"
	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()

	uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
	uiAPI1.HandleFunc("/", UIAPI1RootHandler).Methods("GET")

	uiAPI1.HandleFunc("/blueprints", UIAPI1BlueprintsHandler).Methods("GET")
	uiAPI1.HandleFunc("/blueprints", UIAPI1BlueprintsCreateHandler).Methods("POST")
	uiAPI1.HandleFunc("/blueprints/{id}", UIAPI1BlueprintsDeleteHandler).Methods("DELETE")
	uiAPI1.HandleFunc("/blueprints/{id}", UIAPI1BlueprintsIDHandler).Methods("GET")

	ui := r.PathPrefix("/ui").Subrouter().StrictSlash(true)
	ui.PathPrefix("/{uiDir:(app)|(assets)|(dist)|(node_modules)}").Handler(
		http.StripPrefix("/ui", NoDirListing(
			http.FileServer(http.Dir(viper.GetString("UIRoot"))))))
	ui.HandleFunc("/{_:.*}", UIIndexHandler)

	r.HandleFunc("/{_:.*}", UIIndexHandler)

	go func() {
		Watch(viper.GetInt("WatchPeriod"))
	}()

	ip := viper.GetString("ServerIP")
	port := viper.GetString("ServerPort")
	addr := ip + ":" + port
	log.WithField("address", addr).Infoln("#http", "Listen And Serve")
	log.WithField("UIRoot", viper.GetString("UIRoot")).Infoln("#viper")
	http.ListenAndServe(addr, r)
}

// NoDirListing returns 404 instead of directory listing with http.FileServer
func NoDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// UIAPI1RootHandler returns UI API version
func UIAPI1RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugln("#http,#UIAPI1RootHandler")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"version": 1})
}

// UIAPI1BlueprintsHandler returns a collection of blueprints
func UIAPI1BlueprintsHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugln("#http,#UIAPI1BlueprintsHandler")
	w.Header().Set("Content-Type", "application/vnd.api+json")

	var data []interface{}
	err := store.LoadBucket([]byte(blueprint.DBBucket), &data)

	ans := map[string]interface{}{
		"data": data,
	}

	if err != nil {
		ans["errors"] = map[string]interface{}{
			"details": err,
		}
	}

	json.NewEncoder(w).Encode(ans)
}

// UIAPI1BlueprintsCreateHandler creates and stores a blueprint
func UIAPI1BlueprintsCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	// TODO: ParseFile, then RunBlueprint?

	ans := map[string]interface{}{
		"data": map[string]interface{}{
			"id": "id1",
		},
	}

	w.Header().Set("Location", "/ui/api/v1/blueprints/id1")
	w.WriteHeader(201)

	log.WithFields(log.Fields{
		"request":  r,
		"response": ans,
	}).Debugln("#http,#UIAPI1BlueprintsCreateHandler")
	json.NewEncoder(w).Encode(ans)
}

// UIAPI1BlueprintsDeleteHandler deletes blueprint by ID
func UIAPI1BlueprintsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/vnd.api+json")
	ans := map[string]interface{}{
		"meta": map[string]interface{}{
			"info": "UIAPI1BlueprintsDeleteHandler",
		},
	}

	// Deletion request has been accepted for processing,
	// but the processing has not been completed by the time the server responds.
	// So, Response 202 Accepted
	w.WriteHeader(202)

	log.WithFields(log.Fields{
		"request":  r,
		"response": ans,
		"vars":     vars,
	}).Debugln("#http,#UIAPI1BlueprintsDeleteHandler")
	json.NewEncoder(w).Encode(ans)

	go func() {
		blueprint.DeleteByID(uuid.FromStringOrNil(vars["id"]))
	}()
}

// UIAPI1BlueprintsIDHandler returns a blueprint by id
func UIAPI1BlueprintsIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/vnd.api+json")

	var data interface{}
	err := store.Load([]byte(blueprint.DBBucket), []byte(vars["id"]), &data)

	ans := map[string]interface{}{
		"data": data,
	}

	if err != nil {
		ans["errors"] = map[string]interface{}{
			"details": err,
		}
	}

	log.WithFields(log.Fields{
		"request":  r,
		"response": ans,
		"vars":     vars,
	}).Debugln("#http,#UIAPI1BlueprintsIDHandler")
	json.NewEncoder(w).Encode(ans)
}

// UIIndexHandler returns UI index file
func UIIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugln("#http,#UIIndexHandler")
	indexPath := path.Join(viper.GetString("UIRoot"), "index.html")
	http.ServeFile(w, r, indexPath)
}

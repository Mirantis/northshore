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
	"log"
	"net/http"
	"strings"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/Mirantis/northshore/fsm"
	"github.com/Mirantis/northshore/store"
	"github.com/gorilla/mux"
)

func Run(port string) {
	r := mux.NewRouter()

	uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
	uiAPI1.HandleFunc("/", UIAPI1RootHandler).Methods("GET")

	uiAPI1.HandleFunc("/blueprints", UIAPI1BlueprintsHandler).Methods("GET")

	ui := r.PathPrefix("/ui").Subrouter().StrictSlash(true)
	ui.PathPrefix("/{uiDir:(app)|(assets)|(node_modules)}").Handler(http.StripPrefix("/ui", NoDirListing(http.FileServer(http.Dir("ui/")))))
	ui.HandleFunc("/{_:.*}", UIIndexHandler)

	r.HandleFunc("/{_:.*}", UIIndexHandler)

	//Update frequency for watcher in seconds
	period := 3
	go func() {
		fsm.Watch(period)
	}()

	addr := ":"
	addr += port
	log.Printf("Listening at %s", addr)
	log.Print(http.ListenAndServe(addr, r))
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"version": 1})
}

// UIAPI1BlueprintsHandler returns an collection of blueprints
func UIAPI1BlueprintsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	var data []interface{}
	err := store.LoadBucket([]byte(blueprint.DBBucketBlueprints), &data)

	o := map[string]interface{}{
		"data": data,
	}

	if err != nil {
		o["errors"] = map[string]interface{}{
			"details": err,
		}
	}

	json.NewEncoder(w).Encode(o)
}

// UIIndexHandler returns UI index file
func UIIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("#uiIndexHandler")
	http.ServeFile(w, r, "ui/index.html")
}

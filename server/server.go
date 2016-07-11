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

	"github.com/Mirantis/northshore/fsm"
	"github.com/gorilla/mux"
)

func Run(bpPath string) {
	r := mux.NewRouter()

	uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
	uiAPI1.HandleFunc("/", UIAPI1RootHandler).Methods("GET")

	//uiAPI1.HandleFunc("/blueprints", blueprints).Methods("GET", "POST")

	ui := r.PathPrefix("/ui").Subrouter().StrictSlash(true)
	ui.PathPrefix("/{uiDir:(app)|(assets)|(node_modules)}").Handler(http.StripPrefix("/ui", NoDirListing(http.FileServer(http.Dir("ui/")))))
	ui.HandleFunc("/{_:.*}", UIIndexHandler)

	r.HandleFunc("/{_:.*}", UIIndexHandler)

	states := make(chan map[string]string, 3)

	//Update frequency for watcher in seconds
	period := 3
	go func(c chan map[string]string) {
		fsm.Watch(period, c)
	}(states)

	log.Println("Listening at port 8998")
	http.ListenAndServe(":8998", r)
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

// UIIndexHandler returns UI index file
func UIIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("#uiIndexHandler")
	http.ServeFile(w, r, "ui/index.html")
}

//func blueprints(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/vnd.api+json")
//
//	o := map[string]interface{}{
//		"data": []blueprint.BP{
//			bpl,
//		},
//		"meta": map[string]interface{}{
//			"info": "blueprints",
//		},
//	}
//
//	json.NewEncoder(w).Encode(o)
//}

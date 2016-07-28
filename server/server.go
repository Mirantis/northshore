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
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/Mirantis/northshore/store"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/manyminds/api2go/jsonapi"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func Run() {
	r := mux.NewRouter()

	uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
	uiAPI1.HandleFunc("/", UIAPI1RootHandler).Methods("GET")

	uiAPI1.HandleFunc("/blueprints", UIAPI1BlueprintsHandler).Methods("GET")
	uiAPI1.HandleFunc("/blueprints", UIAPI1BlueprintsCreateHandler).Methods("POST")
	uiAPI1.HandleFunc("/blueprints/{id}", UIAPI1BlueprintsDeleteHandler).Methods("DELETE")
	uiAPI1.HandleFunc("/blueprints/{id}", UIAPI1BlueprintsIDHandler).Methods("GET")
	uiAPI1.HandleFunc("/blueprints/{id}", UIAPI1BlueprintsUpdateHandler).Methods("PATCH")

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

	var ans []byte
	var err error
	var bps []blueprint.Blueprint
	if err := store.LoadBucketAsSlice([]byte(blueprint.DBBucket), &bps); err != nil {
		log.Errorln("Error during get BPs from DB: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ans, err = jsonapi.Marshal(bps); err != nil {
		log.Errorln("Error during marshalling bps: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(ans); err != nil {
		log.Errorln("Error during writing responce: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UIAPI1BlueprintsCreateHandler creates and stores a blueprint
func UIAPI1BlueprintsCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	var body []byte
	var err error
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		log.Errorln("Error during read request body: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var bp blueprint.Blueprint
	if err := jsonapi.Unmarshal(body, &bp); err != nil {
		log.Errorln("Error during unmarshalling body: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Debugln("#CreateHandler, #BP", bp)

	//TODO Handle case if already exists
	store.Save([]byte(blueprint.DBBucket), []byte(bp.GetID()), bp)
	//TODO Automate location handling
	location := "/ui/api/v1/blueprints/" + bp.GetID()

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

// UIAPI1BlueprintsUpdateHandler creates and stores a blueprint
func UIAPI1BlueprintsUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/vnd.api+json")

	var body []byte
	var err error
	var oldBP blueprint.Blueprint
	var newBP blueprint.Blueprint
	if err = store.Load([]byte(blueprint.DBBucket), []byte(vars["id"]), &oldBP); err != nil {
		//TODO Handle error gracefully
		if err.Error() == errors.New("Key does not exist or key is a nested bucket").Error() {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	log.Debugln("#UpateHandler, #OldBP", oldBP)

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		log.Errorln("Error during read request body: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := jsonapi.Unmarshal(body, &newBP); err != nil {
		log.Errorln("Error during unmarshalling body: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Debugln("#UpdateHandler, #NewBP", newBP)
	//TODO Handle case with partial update
	if err := store.Save([]byte(blueprint.DBBucket), []byte(vars["id"]), newBP); err != nil {
		log.Errorln("Error during updating BP in DB: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UIAPI1BlueprintsDeleteHandler deletes blueprint by ID
func UIAPI1BlueprintsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/vnd.api+json")

	// Deletion request has been accepted for processing,
	// but the processing has not been completed by the time the server responds.
	// So, Response 202 Accepted
	w.WriteHeader(http.StatusAccepted)
	log.Debugln("#DeleteHandler", vars["id"])

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

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

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()

		uiAPI1 := r.PathPrefix("/ui/api/v1").Subrouter().StrictSlash(true)
		uiAPI1.HandleFunc("/", uiAPI1RootHandler).Methods("GET")
		uiAPI1.HandleFunc("/action", uiAPI1ActionHandler).Methods("GET", "POST")

		ui := r.PathPrefix("/ui").Subrouter().StrictSlash(true)
		ui.PathPrefix("/{uiDir:(app)|(assets)|(node_modules)}").Handler(http.StripPrefix("/ui", http.FileServer(http.Dir("ui/"))))
		ui.HandleFunc("/{s}", uiIndexHandler)
		ui.HandleFunc("/", uiIndexHandler)

		// with 'nshore run local', you can got to http://localhost:8998/ and see a list of
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

		log.Println("Listening at port 8998")
		http.ListenAndServe(":8998", r)
	},
}

func uiAPI1ActionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Api1ActionHandler"})
}

func uiAPI1RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"version": 1})
}

func uiIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("#uiIndexHandler")
	http.ServeFile(w, r, "ui/index.html")
}

func init() {
	runCmd.AddCommand(localCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// localCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// localCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

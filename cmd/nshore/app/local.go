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
	//"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/gorilla/mux"
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
		r.HandleFunc("/api", apiHandler)
		// with 'nshore run local', you can got to http://localhost:8998/ and see a list of
		// what is in static ... if you put index.html in there, it'll be returned.
		// NB: do not put /static in the path, that'll get you a 404.
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

		log.Println("Listening at port 8998")
		http.ListenAndServe(":8998", r)
	},
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, World! (api)"))
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


// package main
//
// import (
//   "log"
//   "net/http"
//   "github.com/gorilla/mux"
// )
//
// func main() {
//   r := mux.NewRouter()
//   r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("assets/"))))
//
//   log.Println("Listening at port 3000")
//   http.ListenAndServe(":3000", r)
// }

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
	"github.com/Mirantis/northshore/server"
	"github.com/spf13/cobra"
)

var port string

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Run NorthShore local",
	Long:  `Run local HTTP server with BoltDB and watcher for Docker.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(port)
	},
}

func init() {
	localCmd.Flags().StringVarP(&port, "port", "p", "8998", "Port for local server")
	runCmd.AddCommand(localCmd)
}

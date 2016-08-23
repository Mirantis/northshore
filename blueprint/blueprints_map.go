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

package blueprint

import "github.com/Mirantis/northshore/store"

// BlueprintsMap represents store.Storable collection
type BlueprintsMap map[string]*Blueprint

// Bucket implements store.Storable interface
func (items BlueprintsMap) Bucket() []byte {
	return []byte(DBBucket)
}

// Next implements store.Storable interface
func (items *BlueprintsMap) Next(k []byte) interface{} {
	// Check for assignment to entry in nil map
	if *items == nil {
		*items = make(BlueprintsMap)
	}

	item := &Blueprint{}
	(*items)[string(k)] = item
	return (*items)[string(k)]
}

// GetBlueprintsMap returns collection
func GetBlueprintsMap() (items BlueprintsMap, err error) {
	err = store.GetStorable(&items)
	return
}

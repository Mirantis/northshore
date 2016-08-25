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

package store

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"

	"github.com/boltdb/bolt"
)

// Storable defines collection interface
type Storable interface {
	// Bucket returns the bucket name
	Bucket() []byte
	// Next is an iterator, takes key returns ref to item instance
	Next([]byte) interface{}
}

// GetStorable loads all items from boltdb Bucket
func GetStorable(items Storable) error {
	bucket := items.Bucket()
	db := openDBBucket(bucket)
	defer db.Close()

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		b.ForEach(func(k, v []byte) error {

			if err := json.Unmarshal(v, items.Next(k)); err != nil {
				log.Errorln("#DB,#GetStorable", err)
				return err
			}

			return nil
		})

		return nil
	})
	return err
}

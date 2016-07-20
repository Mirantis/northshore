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
	"errors"

	log "github.com/Sirupsen/logrus"

	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
)

func openDBBucket(bucket []byte) *bolt.DB {
	path := viper.GetString("BoltDBPath")
	log.Debugln("Open #DB -> ", path)
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return db
}

// Delete deletes key from boltdb Bucket
func Delete(bucket []byte, key []byte) error {
	db := openDBBucket(bucket)
	defer db.Close()

	if err := db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucket).Delete(key)
	}); err != nil {
		log.Errorln("#DB,#Delete", err)
		return err
	}
	return nil
}

// Load loads item from boltdb Bucket
func Load(bucket []byte, key []byte, v interface{}) error {
	db := openDBBucket(bucket)
	defer db.Close()

	log.Debugln("#DB", "Load data from DB")
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		buf := b.Get(key)
		if buf == nil {
			log.Debugln("#DB,#Load,#Nil")
			return errors.New("Key does not exist or key is a nested bucket")
		}

		if err := json.Unmarshal(buf, &v); err != nil {
			log.Errorln("#DB,#Load", err)
			return err
		}
		return nil
	})
	return err
}

// LoadBucket loads all items from boltdb Bucket
func LoadBucket(bucket []byte, buf *[]interface{}) error {
	db := openDBBucket(bucket)
	defer db.Close()

	log.Debugln("#DB", "Load bucket from DB")
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		b.ForEach(func(k, v []byte) error {
			var item map[string]interface{}
			if err := json.Unmarshal(v, &item); err != nil {
				log.Errorln("#DB,#LoadBucket", err)
				return err
			}
			*buf = append(*buf, item)
			return nil
		})
		return nil
	})
	return err
}

// Save stores item in boltdb Bucket as JSON
func Save(bucket []byte, key []byte, v interface{}) error {
	db := openDBBucket(bucket)
	defer db.Close()

	log.Debugln("#DB", "Save data to DB")
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		vEncoded, err := json.Marshal(v)
		log.Debugln("#DB,#Save,#Encoded", string(vEncoded))
		if err != nil {
			log.Errorln("#DB,#Store", err)
			return err
		}
		if err := b.Put(key, vEncoded); err != nil {
			log.Errorln("#DB,#Store", err)
			return err
		}
		return nil
	})
	return err
}

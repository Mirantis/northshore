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
	"errors"
	"net/http"

	"github.com/Mirantis/northshore/blueprint"
	"github.com/Mirantis/northshore/store"
	log "github.com/Sirupsen/logrus"
	"github.com/manyminds/api2go"
)

// BlueprintResource represents `api2go.CRUD` interface
type BlueprintResource struct{}

// FindAll implements `api2go.FindAll` interface
func (s BlueprintResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	// var bps []blueprint.Blueprint
	// if err := store.LoadBucketAsSlice([]byte(blueprint.DBBucket), &bps); err != nil {

	bps, err := blueprint.LoadAll()
	log.Debugln("#FindAll", bps)

	if err != nil {
		log.Errorln("#BlueprintResource,#FindAll", err)
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}

	return &api2go.Response{Res: bps}, nil
}

// FindOne implements `api2go.CRUD` interface
func (s BlueprintResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	var bp blueprint.Blueprint
	if err := store.Load([]byte(blueprint.DBBucket), []byte(id), &bp); err != nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Errorln("#BlueprintResource,#FindOne", err)
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &api2go.Response{Res: bp}, nil
}

// Delete implements `api2go.CRUD` interface
func (s BlueprintResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	var bp blueprint.Blueprint
	if err := store.Load([]byte(blueprint.DBBucket), []byte(id), &bp); err != nil {
		log.WithFields(log.Fields{
			"id": id,
		}).Errorln("#BlueprintResource,#FindOne", err)
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	go func() {
		bp.Delete()
	}()

	// Deletion request has been accepted for processing,
	// but the processing has not been completed by the time the server responds.
	// So, Response 202 Accepted
	return &api2go.Response{Code: http.StatusAccepted}, nil
}

// Create implements `api2go.CRUD` interface
func (s BlueprintResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	bp, ok := obj.(blueprint.Blueprint)
	if !ok {
		errorMessage := "Invalid instance given"
		log.WithFields(log.Fields{
			"bp":  bp,
			"obj": obj,
		}).Errorln("#BlueprintResource,#Create", errorMessage)
		return &api2go.Response{}, api2go.NewHTTPError(errors.New(errorMessage), errorMessage, http.StatusBadRequest)
	}
	if err := bp.Save(); err != nil {
		log.WithFields(log.Fields{
			"bp":  bp,
			"obj": obj,
		}).Errorln("#BlueprintResource,#Create", err)
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}

	return &api2go.Response{Res: bp, Code: http.StatusCreated}, nil
}

// Update implements `api2go.CRUD` interface
func (s BlueprintResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	bp, ok := obj.(blueprint.Blueprint)
	if !ok {
		errorMessage := "Invalid instance given"
		log.WithFields(log.Fields{
			"bp":  bp,
			"obj": obj,
		}).Errorln("#BlueprintResource,#Update", errorMessage)
		return &api2go.Response{}, api2go.NewHTTPError(errors.New(errorMessage), errorMessage, http.StatusBadRequest)
	}

	// TODO: check for ID from payload is equal to route param
	var bpOld blueprint.Blueprint
	if err := store.Load([]byte(blueprint.DBBucket), []byte(bp.ID.String()), &bpOld); err != nil {
		log.WithFields(log.Fields{
			"id":  bp.ID,
			"obj": obj,
		}).Errorln("#BlueprintResource,#Update", err)
		return &api2go.Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	go func() {
		// TODO:
		// Handle Run/Stop blueprint if bp.State != bpOld.State
		// Handle Create/Delete stages
		// Use stages states from the old one

		if err := bp.Save(); err != nil {
			log.WithFields(log.Fields{
				"bp":  bp,
				"obj": obj,
			}).Errorln("#BlueprintResource,#Update", err)
		}
	}()

	// Update request has been accepted for processing,
	// but the processing has not been completed by the time the server responds.
	// So, Response 202 Accepted
	return &api2go.Response{Code: http.StatusAccepted}, nil
}

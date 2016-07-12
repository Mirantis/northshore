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

import (
	"github.com/Mirantis/northshore/fsm"
	"github.com/Mirantis/northshore/store"
	"github.com/satori/go.uuid"
)

// DBBucketBlueprints defines boltdb bucket for blueprints
const DBBucketBlueprints = "blueprints"

// BP represents a combined data of the Blueprint with States
// TODO: implement LoadBP constructor to instantiate from DB
type BP struct {
	*Blueprint
	*fsm.BlueprintFSM
	UUID uuid.UUID `json:"uuid"`
}

// NewBP creates a BP from Blueprint and stores in DB
func NewBP(blueprint *Blueprint) *BP {
	stagesStates := make(map[string]fsm.StageState)
	for k := range blueprint.Stages {
		stagesStates[k] = fsm.StageStateNew
	}

	bp := &BP{
		blueprint,
		fsm.NewBlueprintFSM(stagesStates),
		uuid.NewV4(),
	}
	store.Store([]byte(DBBucketBlueprints), bp.UUID.Bytes(), bp)
	return bp
}

// Update updates current Blueprint status with Stages
// and stores in DB
func (bp *BP) Update(stagesStates map[string]fsm.StageState) error {
	bp.BlueprintFSM.Update(stagesStates)
	return store.Store([]byte(DBBucketBlueprints), bp.UUID.Bytes(), bp)
}

/*
// NewBPfromJSON creates a BP from JSON
func NewBPfromJSON(data []byte) *BP {

	var buf BP
	if err := json.Unmarshal(data, &buf); err != nil {
		log.Println("#NewBP,#Error", err)
	}

	bp := &BP{
		buf.Blueprint,
		fsm.NewBlueprintFSM(map[string]fsm.StageState(buf.BlueprintFSM)),
		uuid.UUID(buf.UUID),
	}

	log.Println("#NewBP", buf, "\n\n", bp.Blueprint, "\n\n", bp.BlueprintFSM)
	return bp
}
*/

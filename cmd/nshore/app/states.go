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
	"log"

	"github.com/looplab/fsm"
)

type BlueprintState byte
type StageState byte

const (
	BlueprintStateNew BlueprintState = iota
	BlueprintStateProvision
	BlueprintStateActive
	BlueprintStateInactive
)

const (
	StageStateNew StageState = iota
	StageStateCreated
	StageStateRunning
	StageStatePaused
	StageStateStoped
	StageStateDeleted
)

func (state BlueprintState) String() string {
	states := []string{
		"new",
		"provision",
		"active",
		"inactive",
	}
	return states[state]
}

func (state StageState) String() string {
	states := []string{
		"new",
		"created",
		"running",
		"paused",
		"stoped",
		"deleted",
	}
	return states[state]
}

type BlueprintPipeline struct {
	State        BlueprintState
	StagesStates map[string]StageState
	fSM          *fsm.FSM
}

func NewBlueprintPipeline(stages []string) *BlueprintPipeline {
	plStages := map[string]StageState{}
	for _, v := range stages {
		plStages[v] = StageStateNew
	}

	pl := &BlueprintPipeline{
		BlueprintStateNew,
		plStages,
		nil,
	}

	pl.fSM = fsm.NewFSM(
		"new",
		fsm.Events{
			{
				Name: "activate",
				Src:  []string{"inactive", "provision"},
				Dst:  "active",
			},
			{
				Name: "inactivate",
				Src:  []string{"active", "provision"},
				Dst:  "inactive",
			},
			{
				Name: "start",
				Src:  []string{"new"},
				Dst:  "provision",
			},
		},
		fsm.Callbacks{
			"after_event": func(e *fsm.Event) { pl.afterEvent(e) },
			"activate":    func(e *fsm.Event) { pl.afterActivate(e) },
			"inactivate":  func(e *fsm.Event) { pl.afterInactivate(e) },
			"start":       func(e *fsm.Event) { pl.afterStart(e) },
		},
	)

	return pl
}

func (pl *BlueprintPipeline) afterEvent(e *fsm.Event) {
	log.Printf("#BlueprintPipeline,#afterEvent %+v %+v", e, pl)
}

func (pl *BlueprintPipeline) afterActivate(e *fsm.Event) {
	stagesStates := e.Args[0].(map[string]StageState)
	for k, v := range stagesStates {
		pl.StagesStates[k] = v
	}

	for _, v := range pl.StagesStates {
		if v != StageStateRunning {
			e.Cancel()
			return
		}
	}
	pl.State = BlueprintStateActive
}

func (pl *BlueprintPipeline) afterInactivate(e *fsm.Event) {
	stagesStates := e.Args[0].(map[string]StageState)
	for k, v := range stagesStates {
		pl.StagesStates[k] = v
	}
	pl.State = BlueprintStateInactive
}

func (pl *BlueprintPipeline) afterStart(e *fsm.Event) {
	for stage := range pl.StagesStates {
		log.Printf("#BlueprintPipeline,#afterStart Create stage %s", stage)
		// Call to DockerEng
		pl.StagesStates[stage] = StageStateCreated
	}
	pl.State = BlueprintStateProvision
}

func (pl *BlueprintPipeline) Start() {
	pl.fSM.Event("start")
}

func (pl *BlueprintPipeline) Update(stagesStates map[string]StageState) {
	for _, v := range stagesStates {
		switch v {
		case
			StageStatePaused,
			StageStateStoped,
			StageStateDeleted:
			pl.fSM.Event("inactivate", stagesStates)
			return
		}
	}
	pl.fSM.Event("activate", stagesStates)
}

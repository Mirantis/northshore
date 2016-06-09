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

// BlueprintState represents a state of the Blueprint
type BlueprintState byte

// StageState represents a state of the Stage
type StageState byte

const (
	// BlueprintStateNew is default state of the Blueprint
	BlueprintStateNew BlueprintState = iota
	// BlueprintStateProvision is the Blueprint status while provisioning
	BlueprintStateProvision
	// BlueprintStateActive is the Blueprint status when all Stages are up and ready
	BlueprintStateActive
	// BlueprintStateInactive is the Blueprint status when some Stage is down
	BlueprintStateInactive
)

const (
	// StageStateNew is default state of the Stage
	StageStateNew StageState = iota
	// StageStateCreated indicates that container is created
	StageStateCreated
	// StageStateRunning indicates that container is running
	StageStateRunning
	// StageStatePaused indicates that container is paused
	StageStatePaused
	// StageStateStoped indicates that container is stoped
	StageStateStoped
	// StageStateDeleted indicates that container is deleted
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

// BlueprintPipeline represents a Blueprint Pipeline
type BlueprintPipeline struct {
	// State is current Pipeline status
	State BlueprintState
	// StagesStates represents statuses of Pipeline Stages
	StagesStates map[string]StageState
	// fSM is the finite state machine of Pipeline
	fSM *fsm.FSM
}

// NewBlueprintPipeline constructs a Blueprint Pipeline with Stages
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

	// https://godoc.org/github.com/looplab/fsm#NewFSM
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
	pl.State = BlueprintStateProvision
}

// Start creates and runs Stages in Blueprint Pipeline
func (pl *BlueprintPipeline) Start() {
	pl.fSM.Event("start")
}

// Update updates current Blueprint Pipeline status with Stages
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

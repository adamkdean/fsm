//  _____ ____  __  __
// |  ___/ ___||  \/  |
// | |_  \___ \| |\/| |
// |  _|  ___) | |  | |
// |_|   |____/|_|  |_|
//
// Finite State Machine
// (c) 2018 Adam K Dean

package main

import (
	"fmt"
	"github.com/adamkdean/fsm/pkg/fsm"
)

func main() {
	// Create a state map
	sm := map[string][]string{
		"IDLE":     []string{"STARTED"},
		"STARTED":  []string{"STOPPED"},
		"STOPPED":  []string{"STARTED", "FINISHED"},
		"FINISHED": {},
	}

	// Create a new state machine
	fsm := fsm.New()
	fsm.Initialize(sm, "IDLE")

	// Perform a valid transition
	fmt.Printf("CurrentState: %s \n", fsm.CurrentState)
	if err := fsm.Transition("STARTED"); err != nil {
		fmt.Printf("Error transitioning: %v \n", err)
		return
	}
	fmt.Printf("State changed to: %s \n", fsm.CurrentState)

	// Perform an invalid transition
	fmt.Printf("CurrentState: %s \n", fsm.CurrentState)
	if err := fsm.Transition("IDLE"); err != nil {
		fmt.Printf("Error transitioning: %v \n", err)
		return
	}
	fmt.Printf("State changed to: %s \n", fsm.CurrentState)
}

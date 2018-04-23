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
	"time"
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

	// Create a general event
	ch := make(chan string)
	go func() {
		for {
			fmt.Printf("State changed to: %s\n", <-ch)
		}
	}()

	// Hook up the general event to all state transitions
	fsm.OnTransition("IDLE", ch)
	fsm.OnTransition("STARTED", ch)
	fsm.OnTransition("STOPPED", ch)
	fsm.OnTransition("FINISHED", ch)

	// Perform some valid & invalid transitions
	testTransition(fsm, "STARTED")  // Valid
	testTransition(fsm, "STOPPED")  // Valid
	testTransition(fsm, "IDLE")     // Invalid
	testTransition(fsm, "STARTED")  // Valid
	testTransition(fsm, "STOPPED")  // Valid
	testTransition(fsm, "FINISHED") // Valid
	testTransition(fsm, "IDLE")     // Invalid

	// Keep alive
	fmt.Scanln()
}

func testTransition(fsm *fsm.FSM, to string) {
	fmt.Printf("%s -> %s\n", fsm.CurrentState, to)
	if err := fsm.Transition(to); err != nil {
		fmt.Printf("Error transitioning from %s -> %s: %v\n", fsm.CurrentState, to, err)
	}
	time.Sleep(1 * time.Second)
}

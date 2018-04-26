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
	done := make(chan bool)

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

	// Create and hook up the general event to all state transitions
	ch := make(chan string)
	go func() {
		for {
			fmt.Printf("State changed to: %s\n", <-ch)
		}
	}()
	fsm.OnTransition("*", ch)

	// Create and hook up an event for just FINISHED state transition
	chf := make(chan string)
	go func() {
		for {
			<-chf
			fmt.Println("FINISHED! Exiting program...")
			done <- true
		}
	}()
	fsm.OnTransition("FINISHED", chf)

	// Perform some valid & invalid transitions
	testTransition(fsm, "STARTED")  // Valid
	testTransition(fsm, "STOPPED")  // Valid
	testTransition(fsm, "IDLE")     // Invalid
	testTransition(fsm, "STARTED")  // Valid
	testTransition(fsm, "STOPPED")  // Valid
	testTransition(fsm, "FINISHED") // Valid

	// wait for done signal
	<-done
}

func testTransition(fsm *fsm.FSM, to string) {
	fmt.Printf("%s -> %s\n", fsm.CurrentState, to)
	if err := fsm.Transition(to); err != nil {
		fmt.Printf("Error transitioning from %s -> %s: %v\n", fsm.CurrentState, to, err)
	}
	time.Sleep(1 * time.Second)
}

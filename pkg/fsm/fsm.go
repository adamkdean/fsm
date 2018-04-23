//  _____ ____  __  __
// |  ___/ ___||  \/  |
// | |_  \___ \| |\/| |
// |  _|  ___) | |  | |
// |_|   |____/|_|  |_|
//
// Finite State Machine
// (c) 2018 Adam K Dean

package fsm

import (
	"fmt"
	"github.com/thoas/go-funk"
)

// FSM is the finite state machine struct
type FSM struct {
	CurrentState string
	StateMap     map[string][]string
	States       []string
}

// Initialize takes a statemap & initial state
// and initializes the state machine
func (f *FSM) Initialize(sm map[string][]string, s string) {
	f.CurrentState = s
	f.StateMap = sm
	f.States = funk.Keys(sm).([]string)
}

// Transition changes the state when permissable
func (f *FSM) Transition(to string) error {
	if !funk.Contains(f.States, to) {
		return fmt.Errorf("Invalid state: %v", to)
	}

	for _, s := range f.StateMap[f.CurrentState] {
		if s == to {
			f.CurrentState = to
			return nil
		}
	}

	return fmt.Errorf("Invalid transition: %v", to)
}

// New returns a new, empty instance
func New() *FSM {
	return &FSM{}
}

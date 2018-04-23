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
	States       []string
	StateMap     map[string][]string
	EventMap     map[string][]chan string
}

// Initialize takes a state map and an initial
// state and initializes the state machine
func (f *FSM) Initialize(sm map[string][]string, s string) {
	f.CurrentState = s
	f.States = funk.Keys(sm).([]string)
	f.StateMap = sm
	f.EventMap = map[string][]chan string{}
}

// Transition changes the state when permissable
func (f *FSM) Transition(to string) error {
	if err := f.assureStateExists(to); err != nil {
		return err
	}

	// Iterate through all valid transitions and ensure
	// the request transition state is allowed
	for _, s := range f.StateMap[f.CurrentState] {
		if s == to {
			f.CurrentState = to

			// Iterate through events for this new state
			for _, e := range f.EventMap[f.CurrentState] {
				e <- f.CurrentState
			}

			return nil
		}
	}

	return fmt.Errorf("Invalid transition: %v", to)
}

// OnTransition hooks up event channels to state transitions
func (f *FSM) OnTransition(s string, ch chan string) error {
	if err := f.assureStateExists(s); err != nil {
		return err
	}

	f.EventMap[s] = append(f.EventMap[s], ch)
	return nil
}

func (f *FSM) assureStateExists(s string) error {
	if !funk.Contains(f.States, s) {
		return fmt.Errorf("Invalid state: %v", s)
	}
	return nil
}

// New returns a new, empty instance
func New() *FSM {
	return &FSM{}
}

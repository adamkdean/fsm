//  _____ ____  __  __
// |  ___/ ___||  \/  |
// | |_  \___ \| |\/| |
// |  _|  ___) | |  | |
// |_|   |____/|_|  |_|
//
// Finite State Machine
// (c) 2018 Adam K Dean

package example1

import "github.com/adamkdean/fsm/pkg/fsm"

func main() {
	// Create a state map which allows a linear progression of A -> B -> C
	sm := map[string][]string{
		"A": []string{"B"},
		"B": []string{"C"},
		"C": {},
	}

	// Create a new state machine and initialize it to A
	fsm := fsm.New()
	fsm.Initialize(sm, "A")

	// Perform some basic transitions
	fsm.Transition("B") // A -> B
	fsm.Transition("C") // A -> B
}

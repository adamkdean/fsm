//  _____ ____  __  __
// |  ___/ ___||  \/  |
// | |_  \___ \| |\/| |
// |  _|  ___) | |  | |
// |_|   |____/|_|  |_|
//
// Finite State Machine
// (c) 2018 Adam K Dean

package example2

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
	testTransition(fsm, "B") // A -> B = Valid
	testTransition(fsm, "C") // B -> C = Valid
	testTransition(fsm, "A") // C -> A = Invalid
}

func testTransition(fsm *fsm.FSM, to string) {
	fmt.Printf("%s -> %s\n", fsm.CurrentState, to)
	if err := fsm.Transition(to); err != nil {
		fmt.Printf("Error transitioning from %s -> %s: %v\n", fsm.CurrentState, to, err)
	}
}

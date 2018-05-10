```
#  _____ ____  __  __
# |  ___/ ___||  \/  |
# | |_  \___ \| |\/| |
# |  _|  ___) | |  | |
# |_|   |____/|_|  |_|
#
# Finite State Machine
# (c) 2018 Adam K Dean
```

# Finite State Machine

`FiniteStateMachine` is a simple golang implementation of a finite state machine.

## Getting started

### Install

`go get github.com/adamkdean/fsm/pkg/fsm`

### Import

`import "github.com/adamkdean/fsm/pkg/fsm"`

## Examples

The following examples highlight all the current features of this very simple yet powerful implementation.

### Simple

This simple example shows the use of a state machine that allows `A` `->` `B` ``->`` `C`.

```golang
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
```

### Error handling

This example shows the use of error handling resulting from invalid transitions.

```golang
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

func testTransition(fsm *fsm.FiniteStateMachine, to string) {
	fmt.Printf("%s -> %s\n", fsm.CurrentState, to)
	if err := fsm.Transition(to); err != nil {
		fmt.Printf("Error transitioning from %s -> %s: %v\n", fsm.CurrentState, to, err)
	}
}
```

### Events

This example shows the use of `fsm.OnTransition(state string, channel chan string)` to hook up events to transitions.

```golang
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
			done<-true
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

func testTransition(fsm *fsm.FiniteStateMachine, to string) {
	fmt.Printf("%s -> %s\n", fsm.CurrentState, to)
	if err := fsm.Transition(to); err != nil {
		fmt.Printf("Error transitioning from %s -> %s: %v\n", fsm.CurrentState, to, err)
	}
	time.Sleep(1 * time.Second)
}
```

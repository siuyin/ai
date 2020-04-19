package search

import "fmt"

type transitionModel struct{}

// tmS OMIT
// NextState returns the next state give current state s, and action a.
func (tm transitionModel) NextState(s State, a Action) State {
	switch a.ID {
	case 1:
		return State{ID: s.ID - 1}
	case 2:
		return State{ID: s.ID + 1}
	}
	return State{}
}

// tmE OMIT

// aaS OMIT
type availableActions struct{}

func (aa availableActions) Actions(s State) []Action {
	switch {
	case s.ID > 10:
		return []Action{
			Action{ID: 1, Name: "<--"},
		}
	case s.ID < -10:
		return []Action{
			Action{ID: 2, Name: "-->"},
		}
	}
	return []Action{
		Action{ID: 1, Name: "<--"},
		Action{ID: 2, Name: "-->"},
	}
}

// aaE OMIT

func ExampleSearch() {
	// invS OMIT
	start := State{ID: 12}
	goal := State{ID: 4}

	tm := transitionModel{}
	aa := availableActions{}

	g, err := Search(goal, start, tm, aa) // HL01
	if err != nil {
		fmt.Println(err)
	}
	// invE OMIT
	fmt.Println(g.Path())
	// Output:
	// Should fail to demonstrate output

}

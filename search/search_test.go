package search

import "fmt"

// comments ending with OMIT and HL are markers used by the go present program.

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
	// [(4: <--) (5: <--) (6: <--) (7: <--) (8: <--) (9: <--) (10: <--) (11: <--) (12: )]

}

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

func ExampleSearchDFS() {

	start := State{ID: 12}
	goal := State{ID: 4}

	tm := transitionModel{}
	aa := availableActions{}

	g, err := SearchDFS(goal, start, tm, aa) // HL01
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g.Path())
	// Output:
	// [(4: <--) (5: <--) (6: <--) (7: <--) (8: <--) (9: <--) (10: <--) (11: <--) (12: )]
}

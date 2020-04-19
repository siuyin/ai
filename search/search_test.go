package Search

import "fmt"

type transitionModel struct{}

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

func ExampleSearch() {
	start := State{ID: 12}
	goal := State{ID: 4}

	tm := transitionModel{}
	aa := availableActions{}

	g, err := Search(goal, start, tm, aa)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g.Path())
	// Output:
	// ger

}

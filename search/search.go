// Package Search
//  Given a problem with a defined starting point,
//  find a series of actions
//  that will lead to a, preferably optimal, solution.
package Search

import "fmt"

// Search returns the goal state and nil error when successful.
// goal.Path can then be called to provide the path from the
// goal state to the start state.
//
// If Search fails, error is non-nil.
func Search(goal, s State, tm NextStateter, aa Actionsner) (State, error) {
	searcher := newBreadthFirstSearch(goal, s, tm, aa)
	return searcher.search(s)
}

type NextStateter interface {
	NextState(s State, a Action) State
}

type Actionsner interface {
	Actions(s State) []Action
}

type State struct {
	ID           int
	Description  string
	ParentState  *State
	ParentAction Action
}

func (s State) Path() []*State {
	ss := []*State{}
	ss = append(ss, &s)
	for prev := s.ParentState; prev != nil; prev = prev.ParentState {
		ss = append(ss, prev)
	}
	return ss
}
func (s State) String() string {
	return fmt.Sprintf("(%v: %s)", s.ID, s.ParentAction.Name)
}

type Action struct {
	ID   int
	Name string
}

func newBreadthFirstSearch(goal, s State, tm NextStateter, aa Actionsner) *breadthFirstSearch {
	bfs := &breadthFirstSearch{
		Goal:             goal,
		StartState:       s,
		transitionModel:  tm,
		availableActions: aa,
	}

	bfs.q = []State{}
	bfs.discovered = map[State]struct{}{}

	return bfs
}

type breadthFirstSearch struct {
	Goal       State
	StartState State

	q                []State
	discovered       map[State]struct{}
	transitionModel  NextStateter
	availableActions Actionsner
}

// Pseudocode from wikipedia below, where start_v is
// the start vertex or start state.
//  1  procedure BFS(G, start_v) is
//  2      let Q be a queue
//  3      label start_v as discovered
//  4      Q.enqueue(start_v)
//  5      while Q is not empty do
//  6          v := Q.dequeue()
//  7          if v is the goal then
//  8              return v
//  9          for all edges from v to w in G.adjacentEdges(v) do
//  10             if w is not labeled as discovered then
//  11                 label w as discovered
//  12                 w.parent := v
//  13                 Q.enqueue(w)
func (b *breadthFirstSearch) search(startV State) (State, error) {
	b.markDiscovered(startV)
	b.enqueue(startV)
	for b.qLength() > 0 {
		v := b.dequeue()
		if b.atGoal(v) {
			return v, nil
		}
		for _, action := range b.availableActions.Actions(v) {
			w := b.transitionModel.NextState(v, action)
			if !b.isDiscovered(w) {
				b.markDiscovered(w)
			}
			w.ParentState = &v
			w.ParentAction = action
			b.enqueue(w)
		}
	}
	return State{}, fmt.Errorf("search failed to find goal")
}

func (b *breadthFirstSearch) markDiscovered(s State) {
	b.discovered[s] = struct{}{}
}

func (b *breadthFirstSearch) enqueue(s State) {
	b.q = append(b.q, s)
}

func (b *breadthFirstSearch) qLength() int {
	return len(b.q)
}

func (b *breadthFirstSearch) dequeue() State {
	s := b.q[0]
	b.q = b.q[1:]
	return s
}
func (b *breadthFirstSearch) atGoal(s State) bool {
	return s.ID == b.Goal.ID
}

func (b *breadthFirstSearch) isDiscovered(s State) bool {
	_, disc := b.discovered[s]
	return disc
}

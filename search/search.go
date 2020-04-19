// Package search -- given a problem with a defined starting point.
// Find a series of actions
// that will lead to a, preferably optimal, solution.
package search

import "fmt"

// comments ending with OMIT and HL are markers used by the go present program.

// schS OMIT

// Search returns the goal state and nil error when successful.
// goal.Path can then be called to provide the path from the
// goal state tracking back to the start state.
//
// If Search fails, error is non-nil.
func Search(goal, s State, tm NextStateter, aa Actionsner) (State, error) {
	searcher := newBreadthFirstSearch(goal, s, tm, aa) // HL01
	return searcher.search(s)
}

// schE OMIT

// NextStateter is an interface for a transition model that calls NextState.
type NextStateter interface {
	NextState(s State, a Action) State
}

// Actionsner is an interface for an available actions model that calls Actions.
type Actionsner interface {
	Actions(s State) []Action
}

// State encapsulates the state of the world.
// The world is the environment in which the agent operates.
type State struct {
	ID           int
	Description  string
	ParentState  *State
	ParentAction Action
}

// Path retuns path from goal state tracking back to the start state.
func (s State) Path() []*State {
	ss := []*State{}
	ss = append(ss, &s)
	for prev := s.ParentState; prev != nil; prev = prev.ParentState {
		ss = append(ss, prev)
	}
	return ss
}

// strS OMIT

// String returns a string representation of a State.
func (s State) String() string {
	return fmt.Sprintf("(%v: %s)", s.ID, s.ParentAction.Name)
}

// strE OMIT

// Action defines an action that can be taken by the agent.
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
// bfsS OMIT
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
				w.ParentState = &v
				w.ParentAction = action
				b.enqueue(w)
			}
		}
	}
	return State{}, fmt.Errorf("search failed to find goal")
}

// bfsE OMIT

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

// SearchDFS is like Search but searches depth-first instead of breadth-first.
func SearchDFS(goal, s State, tm NextStateter, aa Actionsner) (State, error) {
	searcher := newDepthFirstSearch(goal, s, tm, aa) // HL01
	return searcher.search(s)
}

func newDepthFirstSearch(goal, s State, tm NextStateter, aa Actionsner) *depthFirstSearch {
	dfs := &depthFirstSearch{}

	dfs.Goal = goal
	dfs.StartState = s
	dfs.transitionModel = tm
	dfs.availableActions = aa

	dfs.stack = []State{}
	dfs.discovered = map[State]struct{}{}

	return dfs
}

type depthFirstSearch struct {
	breadthFirstSearch // embeds breadth first search methods. i.e. dfs contains bfs

	stack []State
}

func (d *depthFirstSearch) search(startV State) (State, error) {
	d.markDiscovered(startV)
	d.push(startV)
	for d.sLength() > 0 {
		v := d.pop()
		if d.atGoal(v) {
			return v, nil
		}
		for _, action := range d.availableActions.Actions(v) {
			w := d.transitionModel.NextState(v, action)
			if !d.isDiscovered(w) {
				d.markDiscovered(w)
				w.ParentState = &v
				w.ParentAction = action
				d.push(w)
			}
		}
	}
	return State{}, fmt.Errorf("search failed to find goal")
}

func (d *depthFirstSearch) push(s State) {
	d.stack = append(d.stack, s)
}

func (d *depthFirstSearch) sLength() int {
	return len(d.stack)
}

func (d *depthFirstSearch) pop() State {
	s := d.stack[len(d.stack)-1]
	d.stack = d.stack[:len(d.stack)-1]
	return s
}

// Package know provides methods and functions dealing with knowledge representation.
// The types in this package satisfy the Proposition interface.
package know

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Sentence struct{}

func (s Sentence) Evaluate(model symbolSet) (bool, error) {
	return false, errors.New("nothing to evaluate")
}
func (s Sentence) String() string {
	return ""
}
func (s Sentence) Symbols() symbolSet {
	return symbolSet{}
}

type symbolSet map[Prop]bool

// Prop is short for proposition. Propositions allow evaluation against a model (set of Symbols),
// extracting its constituent set of Symbols and provides
// a human readable string representation.
type Prop interface {
	Evaluate(model symbolSet) (bool, error)
	Symbols() symbolSet
	String() string
}

type Symbol struct {
	Name string
}

func (s Symbol) Evaluate(model symbolSet) (bool, error) {
	val, ok := model[s]
	if !ok {
		return false, fmt.Errorf("variable %v not in model", s)
	}
	return val, nil
}
func (s Symbol) String() string {
	return s.Name
}
func (s Symbol) Symbols() symbolSet {
	return symbolSet{s: false}
}

type Not struct {
	operand Prop
}

func (n Not) Evaluate(model symbolSet) (bool, error) {
	val, err := n.operand.Evaluate(model)
	return !val, err
}
func (n Not) String() string {
	return fmt.Sprintf("¬%s", parenthesize(n.operand.String()))
}

func parenthesize(s string) string {
	if blank(s) || isAlphaN(s) || (parenthesized(s) && balancedParens(s)) {
		return s
	}
	return fmt.Sprintf("(%s)", s)
}
func blank(s string) bool {
	return len(s) == 0
}
func isAlphaN(s string) bool {
	return unicode.IsLetter(rune(s[0]))
}
func parenthesized(s string) bool {
	return s[0:1] == "(" && s[len(s)-1:] == ")"
}
func balancedParens(s string) bool {
	count := 0
	for char := range s {
		switch char {
		case '(':
			count++
		case ')':
			count--
		}
	}
	return count == 0
}

func (n Not) Symbols() symbolSet {
	return n.operand.Symbols()
}

type And struct {
	// A conjunct is a term in an AND clause. Eg. {A,B} are conjuncts in "A AND B".)akjjjjj
	conjuncts []Prop
}

func (a And) Add(conjunct Prop) {
	a.conjuncts = append(a.conjuncts, conjunct)
}
func (a And) Evaluate(model symbolSet) (bool, error) {
	for _, elem := range a.conjuncts {
		val, err := elem.Evaluate(model)
		if err != nil {
			return false, err
		}
		if val == false {
			return false, nil
		}
	}
	return true, nil
}
func (a And) String() string {
	symbs := symbolStrings(a.conjuncts)
	return strings.Join(symbs, " ^ ")
}
func symbolStrings(e []Prop) []string {
	symbs := []string{}
	for _, elem := range e {
		for symb := range elem.Symbols() {
			symbs = append(symbs, symb.String())
		}
	}
	return symbs
}
func (a And) Symbols() symbolSet {
	return symbolsSet(a.conjuncts)
}
func symbolsSet(e []Prop) symbolSet {
	symbs := symbolSet{}
	for _, elem := range e {
		for symb := range elem.Symbols() {
			symbs[symb] = false
		}
	}
	return symbs
}

type Or struct {
	disjuncts []Prop
}

func (o Or) Add(disjunct Prop) {
	o.disjuncts = append(o.disjuncts, disjunct)
}
func (o Or) Evaluate(model symbolSet) (bool, error) {
	for _, elem := range o.disjuncts {
		val, err := elem.Evaluate(model)
		if err != nil {
			return false, err
		}
		if val == true {
			return true, nil
		}
	}
	return false, nil
}
func (o Or) String() string {
	symbs := symbolStrings(o.disjuncts)
	return strings.Join(symbs, " v ")
}
func (o Or) Symbols() symbolSet {
	return symbolsSet(o.disjuncts)
}

type Implication struct {
	a, b Prop
}

func (i Implication) Evaluate(model symbolSet) (bool, error) {
	return Or{
		[]Prop{
			Not{i.a},
			i.b},
	}.Evaluate(model)
}
func (i Implication) String() string {
	return fmt.Sprintf("%s => %s", i.a, i.b)
}
func (i Implication) Symbols() symbolSet {
	return union(i.a.Symbols(), i.b.Symbols())
}
func union(a, b symbolSet) symbolSet {
	s := symbolSet{}
	for as := range a {
		s[as] = a[as]
	}
	for bs := range b {
		s[bs] = b[bs]
	}
	return s
}

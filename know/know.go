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

func Sym(n string) KSymbol {
	return KSymbol{Name: n}
}

type KSymbol struct {
	Name string
}

func (s KSymbol) Evaluate(model symbolSet) (bool, error) {
	val, ok := model[s]
	if !ok {
		return false, fmt.Errorf("variable %v not in model", s)
	}
	return val, nil
}
func (s KSymbol) String() string {
	return s.Name
}
func (s KSymbol) Symbols() symbolSet {
	return symbolSet{s: false}
}

func Not(op Prop) KNot {
	return KNot{Operand: op}
}

type KNot struct {
	Operand Prop
}

func (n KNot) Evaluate(model symbolSet) (bool, error) {
	val, err := n.Operand.Evaluate(model)
	return !val, err
}
func (n KNot) String() string {
	return fmt.Sprintf("¬%s", parenthesize(n.Operand.String()))
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

func (n KNot) Symbols() symbolSet {
	return n.Operand.Symbols()
}

func And(a ...Prop) KAnd {
	r := KAnd{}
	r.Conjuncts = append(r.Conjuncts, a...)
	return r
}

type KAnd struct {
	// A conjunct is a term in an AND clause. Eg. {A,B} are Conjuncts in "A AND B".)akjjjjj
	Conjuncts []Prop
}

func (a KAnd) Add(conjunct Prop) {
	a.Conjuncts = append(a.Conjuncts, conjunct)
}
func (a KAnd) Evaluate(model symbolSet) (bool, error) {
	for _, elem := range a.Conjuncts {
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
func (a KAnd) String() string {
	symbs := symbolStrings(a.Conjuncts)
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
func (a KAnd) Symbols() symbolSet {
	return symbolsSet(a.Conjuncts)
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

func Or(d ...Prop) Prop {
	r := KOr{}
	r.Disjuncts = append(r.Disjuncts, d...)
	return r
}

type KOr struct {
	Disjuncts []Prop
}

func (o KOr) Add(disjunct Prop) {
	o.Disjuncts = append(o.Disjuncts, disjunct)
}
func (o KOr) Evaluate(model symbolSet) (bool, error) {
	for _, elem := range o.Disjuncts {
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
func (o KOr) String() string {
	symbs := symbolStrings(o.Disjuncts)
	return strings.Join(symbs, " v ")
}
func (o KOr) Symbols() symbolSet {
	return symbolsSet(o.Disjuncts)
}

func Implication(a, b Prop) Prop {
	return KImplication{A: a, B: b}
}

type KImplication struct {
	A, B Prop
}

func (i KImplication) Evaluate(model symbolSet) (bool, error) {
	return KOr{
		[]Prop{
			KNot{i.A},
			i.B},
	}.Evaluate(model)
}
func (i KImplication) String() string {
	return fmt.Sprintf("%s => %s", i.A, i.B)
}
func (i KImplication) Symbols() symbolSet {
	return union(i.A.Symbols(), i.B.Symbols())
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

func Biconditional(a, b Prop) KBiconditional {
	return KBiconditional{A: a, B: b}
}

type KBiconditional struct {
	A, B Prop
}

func (b KBiconditional) Evaluate(model symbolSet) (bool, error) {
	return KOr{[]Prop{
		KAnd{[]Prop{b.A, b.B}},
		KAnd{[]Prop{KNot{b.A}, KNot{b.B}}},
	}}.Evaluate(model)
}
func (b KBiconditional) String() string {
	return fmt.Sprintf("%s <=> %s", b.A, b.B)
}
func (b KBiconditional) Symbols() symbolSet {
	return union(b.A.Symbols(), b.B.Symbols())
}

func ModelCheck(knowledge, query Prop) (bool, error) {
	symbols := union(knowledge.Symbols(), query.Symbols())
	model := symbolSet{}
	return checkAll(knowledge, query, symbols, model)
}
func checkAll(knowledge, query Prop, symbols, model symbolSet) (bool, error) {
	if len(symbols) == 0 {
		kbTrue, err := knowledge.Evaluate(model)
		if err != nil {
			return kbTrue, err
		}
		if kbTrue {
			return query.Evaluate(model)
		}
		return true, nil
	}

	remainingSymbols := copySet(symbols)
	p := pop(remainingSymbols)

	modelTrue, modelFalse := genModelsWithSymbolTrueFalse(model, p)

	vT, err := checkAll(knowledge, query, remainingSymbols, modelTrue)
	if err != nil {
		return vT, err
	}
	vF, err := checkAll(knowledge, query, remainingSymbols, modelFalse)
	if err != nil {
		return vF, err
	}
	return vT && vF, nil
}
func copySet(s symbolSet) symbolSet {
	r := symbolSet{}
	for se := range s {
		r[se] = s[se]
	}
	return r
}
func pop(s symbolSet) Prop {
	for se := range s {
		delete(s, se)
		return se
	}
	return KSymbol{}
}
func genModelsWithSymbolTrueFalse(model symbolSet, p Prop) (symbolSet, symbolSet) {
	modelTrue := copySet(model)
	modelTrue[p] = true

	modelFalse := copySet(model)
	modelFalse[p] = false
	return modelTrue, modelFalse
}

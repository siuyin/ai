package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

	k "github.com/siuyin/ai/know"
)

func main() {
	fmt.Println("Who owns the zebra and who drinks water?")
	//syms := allSymbols()
	//kb := addKnowledge()
	//ents := findEntailments(kb, syms)
	//for i, ent := range ents {
	//	fmt.Printf("%d: %s\n", i, ent)
	//}
}

func and(p symSet, q symSet) symSet {
	r := symSet{}
	for ep := range p {
		if in(q, ep) {
			r[ep] = struct{}{}
		}
	}
	return r
}
func in(p symSet, e k.Prop) bool {
	_, ok := p[e]
	if ok {
		return true
	}
	return false
}
func allSymbols() symSet {
	var mAllSymbols map[k.Prop]struct{}
	if mAllSymbols == nil {
		mAllSymbols = map[k.Prop]struct{}{}
	}
	if len(mAllSymbols) != 0 {
		return mAllSymbols
	}
	syms := map[k.Prop]struct{}{}
	for _, ha := range houseAddress() {
		for _, c := range houseColours() {
			for _, n := range nationalities() {
				for _, cig := range cigBrands() {
					for _, d := range drinks() {
						for _, p := range pets() {
							syms[k.Sym(fmt.Sprintf("%s %s %s %s %s %s", ha, c, n, cig, d, p))] = struct{}{}
						}
					}
				}
			}
		}
	}
	mAllSymbols = syms
	return syms // 5^6 = 15625 symbols!
}

type symSet map[k.Prop]struct{}

func houseAddress() []string {
	return []string{"1", "2", "3", "4", "5"}
}
func houseColours() []string {
	return []string{"Yellow", "Blue", "Red", "Green", "Ivory"}
}
func nationalities() []string {
	return []string{"Norwegian", "Ukranian", "English", "Japanese", "Spanish"}
}
func cigBrands() []string {
	return []string{"Kools", "Chester", "Old Gold", "Parl", "Lucky Strike"}
}
func drinks() []string {
	return []string{"water", "tea", "milk", "coffee", "orange juice"}
}
func pets() []string {
	return []string{"fox", "horse", "snails", "zebra", "dog"}
}

func addKnowledge() k.Prop {
	kb := k.And()

	kb = kb.Add(is("English", "Red"))
	kb = kb.Add(is("Spanish", "dog"))
	kb = kb.Add(is("coffee", "Green"))
	kb = kb.Add(is("Old Gold", "snails"))
	kb = kb.Add(is("Kools", "Yellow"))
	kb = kb.Add(is("milk", "3"))
	kb = kb.Add(is("Norwegian", "1"))
	kb = kb.Add(is("Lucky Strike", "orange juice"))
	kb = kb.Add(is("Japanese", "Parl"))
	return kb
}

// func is("English","Red") and not non("English","Red")
func is(s ...string) k.Prop {
	r := k.Or()
	//er := symbols(contains, "English", "Red")
	er := symbols(contains, s...)
	for sym := range er {
		r.Add(k.Implication(sym, k.Not(allSymbolsMinus(sym))))
	}
	return r
}
func copySet(s symSet) symSet {
	o := symSet{}
	for e := range s {
		o[e] = struct{}{}
	}
	return o
}

func allSymbolsMinus(e k.Prop) k.Prop {
	cpy := copySet(allSymbols())
	delete(cpy, e)

	ret := k.Or()
	for f := range cpy {
		ret.Add(f)
	}
	return ret
}

func contains(s string, args ...string) bool {
	for _, a := range args {
		if !strings.Contains(s, a) {
			return false
		}
	}
	return true
}
func doesNotContain(s string, args ...string) bool {
	for _, a := range args {
		if strings.Contains(s, a) {
			return false
		}
	}
	return true
}
func symbols(fn symFil, args ...string) symSet {
	var memo map[string]symSet
	if memo == nil {
		memo = make(map[string]symSet)
	}

	key, val, found := findMemo(memo, fn, args...)
	if found {
		return val
	}
	memo, syms := filter(memo, key, fn, args...)
	return syms
}

type symFil func(s string, args ...string) bool

func findMemo(memo map[string]symSet, fn symFil, args ...string) (string, symSet, bool) {
	key := fmt.Sprint(funcName(fn), args)
	if v, ok := memo[key]; ok {
		return key, v, true
	}
	return "", symSet{}, false
}
func funcName(i interface{}) string {
	return runtime.FuncForPC(
		reflect.ValueOf(i).Pointer(),
	).Name()
}
func filter(memo map[string]symSet, key string, fn symFil, args ...string) (map[string]symSet, symSet) {
	syms := symSet{}
	for sym := range allSymbols() {
		if fn(sym.String(), args...) {
			syms[sym] = struct{}{}
		}
	}
	memo[key] = syms
	return memo, syms
}

func findEntailments(kb k.Prop, symbs map[k.Prop]struct{}) []k.Prop {
	ents := []k.Prop{}
	for sym := range symbs {
		val, err := k.ModelCheck(kb, sym)
		if err != nil {
			log.Fatalf("findEntailments: %v", err)
		}
		if val {
			ents = append(ents, sym)
		}
	}
	return ents
}

type house struct {
	pos    int
	colour string
}

func (h house) String() string {
	return fmt.Sprintf("House%d:%s", h.pos, h.colour)
}

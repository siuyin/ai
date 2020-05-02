package main

import (
	"fmt"
	"log"

	k "github.com/siuyin/ai/know"
)

func main() {
	fmt.Println("A knight always tells the truth. A knave always lies.")
	fmt.Println(`I said: "I am both a knight and a knave!"`)

	kni := k.Sym("I am a knight!")
	kna := k.Sym("I am a knave.")

	kb := knowledgeAbout(kni, kna) // HL

	ents := findEntailments(kb) // HL
	if len(ents) == 0 {
		fmt.Println("Knowledge base could not make any inferences")
	}
	for _, ent := range ents {
		fmt.Println(ent)
	}
}
func knowledgeAbout(kni, kna k.Prop) k.Prop {
	kb := k.And() // kb is a series of propositions known to be true.
	kb = kb.Add(k.Or(
		k.And(kni, k.Not(kna)),
		k.And(kna, k.Not(kni)),
	))
	kb = kb.Add(k.Implication(kni, k.And(kni, kna)))
	kb = kb.Add(k.Implication(kna, k.Not(k.And(kni, kna))))
	return kb
}
func findEntailments(kb k.Prop) []string {
	entailments := []string{}
	for qry := range kb.Symbols() { // HL
		val, err := k.ModelCheck(kb, qry)
		if err != nil {
			log.Fatal(err)
		}
		if val {
			entailments = append(entailments,
				fmt.Sprintf("\nKnowledge base inferred: %q\n", qry))
		}
	}
	return entailments
}

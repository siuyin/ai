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

	kb := k.And(
		k.Or(
			k.And(kni, k.Not(kna)),
			k.And(kna, k.Not(kni)),
		),
		k.Implication(kni, k.And(kni, kna)),
		k.Implication(kna, k.Not(k.And(kni, kna))),
	)
	for _, qry := range []k.Prop{kni, kna} {
		val, err := k.ModelCheck(kb, qry)
		if err != nil {
			log.Fatal(err)
		}
		if val {
			fmt.Printf("\nKnowledge base inferred: %q\n", qry)
			return
		}
	}
	fmt.Printf("\nKnowledge base could not make any inferences.\n")
}

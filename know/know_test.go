package know

import "testing"

func TestSymbol(t *testing.T) {
	p := Symbol{"P"}
	if p.String() != "P" {
		t.Errorf("unexpected value: %v", p.String())
	}
}
func TestModel(t *testing.T) {
	p := Symbol{"P"}
	m := symbolSet{p: true}
	if m[p] != true {
		t.Error("symbol p should be true")
	}
}
func TestNot(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: Not{Symbol{"P"}}, model: symbolSet{Symbol{"P"}: true, Symbol{"out"}: false}},
		{r: Not{Symbol{"P"}}, model: symbolSet{Symbol{"P"}: false, Symbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "¬P" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		if val, _ := d.r.Evaluate(d.model); val != d.model[Symbol{"out"}] {

			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}
func TestAnd(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: And{[]Prop{Symbol{"P"}, Symbol{"Q"}}},
			model: symbolSet{Symbol{"P"}: true, Symbol{"Q"}: true, Symbol{"out"}: true}},
		{r: And{[]Prop{Symbol{"P"}, Symbol{"Q"}}},
			model: symbolSet{Symbol{"P"}: true, Symbol{"Q"}: false, Symbol{"out"}: false}},
		{r: And{[]Prop{Symbol{"P"}, Symbol{"Q"}}},
			model: symbolSet{Symbol{"P"}: false, Symbol{"Q"}: true, Symbol{"out"}: false}},
		{r: And{[]Prop{Symbol{"P"}, Symbol{"Q"}}},
			model: symbolSet{Symbol{"P"}: false, Symbol{"Q"}: false, Symbol{"out"}: false}},
	}
	for i, d := range dat {
		if d.r.String() != "P ^ Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[Symbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}
func TestOr(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: Or{[]Prop{Symbol{"P"}, Symbol{"Q"}}},
			model: symbolSet{Symbol{"P"}: true, Symbol{"Q"}: true, Symbol{"out"}: true}},
		{r: Or{[]Prop{Symbol{"P"}, Symbol{"Q"}}},
			model: symbolSet{Symbol{"P"}: true, Symbol{"Q"}: false, Symbol{"out"}: true}},
		{r: Or{[]Prop{Symbol{"P"}, Symbol{"Q"}}},
			model: symbolSet{Symbol{"P"}: false, Symbol{"Q"}: true, Symbol{"out"}: true}},
		{r: Or{[]Prop{Symbol{"P"}, Symbol{"Q"}}},
			model: symbolSet{Symbol{"P"}: false, Symbol{"Q"}: false, Symbol{"out"}: false}},
	}
	for i, d := range dat {
		if d.r.String() != "P v Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[Symbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}

func TestAndOld(t *testing.T) {
	dat := []struct {
		a, b SentenceOld
		c    SentenceOld
	}{
		{SentenceOld{"A", true}, SentenceOld{"B", true}, // 0
			SentenceOld{"(A AND B)", true}},
		{SentenceOld{"A", true}, SentenceOld{"B", false}, // 1
			SentenceOld{"(A AND (NOT B))", false}},
	}
	for i, d := range dat {
		r := d.a.And(d.b)
		if r != d.c {
			t.Errorf("case %d: nexpected value: %s, %v", i, r, r.Val)
		}
	}
}

func TestSentenceOld(t *testing.T) {
	q := SentenceOld{"It is sunny", false}
	if q.String() != "(NOT It is sunny)" {
		t.Errorf("unexpected value: %s, %v", q.String(), q.Val)
	}
}

func TestOrOld(t *testing.T) {
	dat := []struct {
		a, b SentenceOld
		c    SentenceOld
	}{
		{SentenceOld{"A", true}, SentenceOld{"B", false}, // 0
			SentenceOld{"(A OR (NOT B))", true}},
		{SentenceOld{"A", false}, SentenceOld{"B", false}, // 1
			SentenceOld{"((NOT A) OR (NOT B))", false}},
	}
	for i, d := range dat {
		r := d.a.Or(d.b)
		if r != d.c {
			t.Errorf("case: %d: unexpected value: %s, %v", i, r, r.Val)
		}
	}
}

func TestImpliesOld(t *testing.T) {
	dat := []struct {
		a, b SentenceOld
		c    SentenceOld
	}{
		{SentenceOld{"A", true}, SentenceOld{"B", true}, // 0
			SentenceOld{"(A IMPLIES B)", true}},
		{SentenceOld{"A", true}, SentenceOld{"B", false}, // 1
			SentenceOld{"(A IMPLIES (NOT B))", false}},
	}
	for i, d := range dat {
		r := d.a.Implies(d.b)
		if r != d.c {
			t.Errorf("case: %d: unexpected value: %s, %v", i, r, r.Val)
		}
	}
}

func TestBiconditionalOld(t *testing.T) {
	dat := []struct {
		a, b SentenceOld
		c    SentenceOld
	}{
		{SentenceOld{"A", true}, SentenceOld{"B", true}, // 0
			SentenceOld{"(A BICONDITIONAL B)", true}},
		{SentenceOld{"A", true}, SentenceOld{"B", false}, // 1
			SentenceOld{"(A BICONDITIONAL (NOT B))", false}},
	}
	for i, d := range dat {
		r := d.a.Biconditional(d.b)
		if r != d.c {
			t.Errorf("case: %d: unexpected value: %s, %v", i, r, r.Val)
		}
	}
}

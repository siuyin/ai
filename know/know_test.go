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

func TestImplication(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: Implication{Symbol{"P"}, Symbol{"Q"}},
			model: symbolSet{Symbol{"P"}: true, Symbol{"Q"}: true, Symbol{"out"}: true}},
		{r: Implication{Symbol{"P"}, Symbol{"Q"}},
			model: symbolSet{Symbol{"P"}: true, Symbol{"Q"}: false, Symbol{"out"}: false}},
		{r: Implication{Symbol{"P"}, Symbol{"Q"}},
			model: symbolSet{Symbol{"P"}: false, Symbol{"Q"}: true, Symbol{"out"}: true}},
		{r: Implication{Symbol{"P"}, Symbol{"Q"}},
			model: symbolSet{Symbol{"P"}: false, Symbol{"Q"}: false, Symbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "P => Q" {
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
func TestBiconditional(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: Biconditional{Symbol{"P"}, Symbol{"Q"}},
			model: symbolSet{Symbol{"P"}: true, Symbol{"Q"}: true, Symbol{"out"}: true}},
		{r: Biconditional{Symbol{"P"}, Symbol{"Q"}},
			model: symbolSet{Symbol{"P"}: true, Symbol{"Q"}: false, Symbol{"out"}: false}},
		{r: Biconditional{Symbol{"P"}, Symbol{"Q"}},
			model: symbolSet{Symbol{"P"}: false, Symbol{"Q"}: true, Symbol{"out"}: false}},
		{r: Biconditional{Symbol{"P"}, Symbol{"Q"}},
			model: symbolSet{Symbol{"P"}: false, Symbol{"Q"}: false, Symbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "P <=> Q" {
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

func TestPop(t *testing.T) {
	p := Symbol{"P"}
	ss := symbolSet{p: false}
	q := pop(ss)
	if q != p {
		t.Errorf("unexpected value: %v", q)
	}
	if len(ss) != 0 {
		t.Errorf("unexpected value: %v", len(ss))
	}
}

func TestModelCheck(t *testing.T) {
	p := Symbol{"P"}
	q := Symbol{"Q"}

	kb := And{[]Prop{p, q}}
	val, err := ModelCheck(kb, q)
	if err != nil {
		t.Error(err)
	}
	if val != true {
		t.Errorf("unexpected value: %v", val)
	}
}

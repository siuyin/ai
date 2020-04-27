package know

import "testing"

func TestSymbol(t *testing.T) {
	p := symbol{"P"}
	if p.String() != "P" {
		t.Errorf("unexpected value: %v", p.String())
	}
}
func TestModel(t *testing.T) {
	p := symbol{"P"}
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
		{r: not{symbol{"P"}}, model: symbolSet{symbol{"P"}: true, symbol{"out"}: false}},
		{r: not{symbol{"P"}}, model: symbolSet{symbol{"P"}: false, symbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "¬P" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		if val, _ := d.r.Evaluate(d.model); val != d.model[symbol{"out"}] {

			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}
func TestAnd(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: and{[]Prop{symbol{"P"}, symbol{"Q"}}},
			model: symbolSet{symbol{"P"}: true, symbol{"Q"}: true, symbol{"out"}: true}},
		{r: and{[]Prop{symbol{"P"}, symbol{"Q"}}},
			model: symbolSet{symbol{"P"}: true, symbol{"Q"}: false, symbol{"out"}: false}},
		{r: and{[]Prop{symbol{"P"}, symbol{"Q"}}},
			model: symbolSet{symbol{"P"}: false, symbol{"Q"}: true, symbol{"out"}: false}},
		{r: and{[]Prop{symbol{"P"}, symbol{"Q"}}},
			model: symbolSet{symbol{"P"}: false, symbol{"Q"}: false, symbol{"out"}: false}},
	}
	for i, d := range dat {
		if d.r.String() != "P ^ Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[symbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}
func TestOr(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: or{[]Prop{symbol{"P"}, symbol{"Q"}}},
			model: symbolSet{symbol{"P"}: true, symbol{"Q"}: true, symbol{"out"}: true}},
		{r: or{[]Prop{symbol{"P"}, symbol{"Q"}}},
			model: symbolSet{symbol{"P"}: true, symbol{"Q"}: false, symbol{"out"}: true}},
		{r: or{[]Prop{symbol{"P"}, symbol{"Q"}}},
			model: symbolSet{symbol{"P"}: false, symbol{"Q"}: true, symbol{"out"}: true}},
		{r: or{[]Prop{symbol{"P"}, symbol{"Q"}}},
			model: symbolSet{symbol{"P"}: false, symbol{"Q"}: false, symbol{"out"}: false}},
	}
	for i, d := range dat {
		if d.r.String() != "P v Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[symbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}

func TestImplication(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: implication{symbol{"P"}, symbol{"Q"}},
			model: symbolSet{symbol{"P"}: true, symbol{"Q"}: true, symbol{"out"}: true}},
		{r: implication{symbol{"P"}, symbol{"Q"}},
			model: symbolSet{symbol{"P"}: true, symbol{"Q"}: false, symbol{"out"}: false}},
		{r: implication{symbol{"P"}, symbol{"Q"}},
			model: symbolSet{symbol{"P"}: false, symbol{"Q"}: true, symbol{"out"}: true}},
		{r: implication{symbol{"P"}, symbol{"Q"}},
			model: symbolSet{symbol{"P"}: false, symbol{"Q"}: false, symbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "P => Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[symbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}
func TestBiconditional(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: biconditional{symbol{"P"}, symbol{"Q"}},
			model: symbolSet{symbol{"P"}: true, symbol{"Q"}: true, symbol{"out"}: true}},
		{r: biconditional{symbol{"P"}, symbol{"Q"}},
			model: symbolSet{symbol{"P"}: true, symbol{"Q"}: false, symbol{"out"}: false}},
		{r: biconditional{symbol{"P"}, symbol{"Q"}},
			model: symbolSet{symbol{"P"}: false, symbol{"Q"}: true, symbol{"out"}: false}},
		{r: biconditional{symbol{"P"}, symbol{"Q"}},
			model: symbolSet{symbol{"P"}: false, symbol{"Q"}: false, symbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "P <=> Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[symbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}

func TestPop(t *testing.T) {
	p := symbol{"P"}
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
	p := symbol{"P"}
	q := symbol{"Q"}

	kb := and{[]Prop{p, q}}
	val, err := ModelCheck(kb, q)
	if err != nil {
		t.Error(err)
	}
	if val != true {
		t.Errorf("unexpected value: %v", val)
	}
}

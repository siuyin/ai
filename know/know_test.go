package know

import "testing"

func TestSymbol(t *testing.T) {
	p := KSymbol{"P"}
	if p.String() != "P" {
		t.Errorf("unexpected value: %v", p.String())
	}
}
func TestModel(t *testing.T) {
	p := KSymbol{"P"}
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
		{r: KNot{KSymbol{"P"}}, model: symbolSet{KSymbol{"P"}: true, KSymbol{"out"}: false}},
		{r: KNot{KSymbol{"P"}}, model: symbolSet{KSymbol{"P"}: false, KSymbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "¬P" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		if val, _ := d.r.Evaluate(d.model); val != d.model[KSymbol{"out"}] {

			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}
func TestAnd(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: KAnd{[]Prop{KSymbol{"P"}, KSymbol{"Q"}}},
			model: symbolSet{KSymbol{"P"}: true, KSymbol{"Q"}: true, KSymbol{"out"}: true}},
		{r: KAnd{[]Prop{KSymbol{"P"}, KSymbol{"Q"}}},
			model: symbolSet{KSymbol{"P"}: true, KSymbol{"Q"}: false, KSymbol{"out"}: false}},
		{r: KAnd{[]Prop{KSymbol{"P"}, KSymbol{"Q"}}},
			model: symbolSet{KSymbol{"P"}: false, KSymbol{"Q"}: true, KSymbol{"out"}: false}},
		{r: KAnd{[]Prop{KSymbol{"P"}, KSymbol{"Q"}}},
			model: symbolSet{KSymbol{"P"}: false, KSymbol{"Q"}: false, KSymbol{"out"}: false}},
	}
	for i, d := range dat {
		if d.r.String() != "P ^ Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[KSymbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}
func TestOr(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: KOr{[]Prop{KSymbol{"P"}, KSymbol{"Q"}}},
			model: symbolSet{KSymbol{"P"}: true, KSymbol{"Q"}: true, KSymbol{"out"}: true}},
		{r: KOr{[]Prop{KSymbol{"P"}, KSymbol{"Q"}}},
			model: symbolSet{KSymbol{"P"}: true, KSymbol{"Q"}: false, KSymbol{"out"}: true}},
		{r: KOr{[]Prop{KSymbol{"P"}, KSymbol{"Q"}}},
			model: symbolSet{KSymbol{"P"}: false, KSymbol{"Q"}: true, KSymbol{"out"}: true}},
		{r: KOr{[]Prop{KSymbol{"P"}, KSymbol{"Q"}}},
			model: symbolSet{KSymbol{"P"}: false, KSymbol{"Q"}: false, KSymbol{"out"}: false}},
	}
	for i, d := range dat {
		if d.r.String() != "P v Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[KSymbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}

func TestImplication(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: KImplication{KSymbol{"P"}, KSymbol{"Q"}},
			model: symbolSet{KSymbol{"P"}: true, KSymbol{"Q"}: true, KSymbol{"out"}: true}},
		{r: KImplication{KSymbol{"P"}, KSymbol{"Q"}},
			model: symbolSet{KSymbol{"P"}: true, KSymbol{"Q"}: false, KSymbol{"out"}: false}},
		{r: KImplication{KSymbol{"P"}, KSymbol{"Q"}},
			model: symbolSet{KSymbol{"P"}: false, KSymbol{"Q"}: true, KSymbol{"out"}: true}},
		{r: KImplication{KSymbol{"P"}, KSymbol{"Q"}},
			model: symbolSet{KSymbol{"P"}: false, KSymbol{"Q"}: false, KSymbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "P => Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[KSymbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}
func TestBiconditional(t *testing.T) {
	dat := []struct {
		r     Prop
		model symbolSet
	}{
		{r: KBiconditional{KSymbol{"P"}, KSymbol{"Q"}},
			model: symbolSet{KSymbol{"P"}: true, KSymbol{"Q"}: true, KSymbol{"out"}: true}},
		{r: KBiconditional{KSymbol{"P"}, KSymbol{"Q"}},
			model: symbolSet{KSymbol{"P"}: true, KSymbol{"Q"}: false, KSymbol{"out"}: false}},
		{r: KBiconditional{KSymbol{"P"}, KSymbol{"Q"}},
			model: symbolSet{KSymbol{"P"}: false, KSymbol{"Q"}: true, KSymbol{"out"}: false}},
		{r: KBiconditional{KSymbol{"P"}, KSymbol{"Q"}},
			model: symbolSet{KSymbol{"P"}: false, KSymbol{"Q"}: false, KSymbol{"out"}: true}},
	}
	for i, d := range dat {
		if d.r.String() != "P <=> Q" {
			t.Errorf("case %d: unexpected value: %v", i, d.r.String())
		}
		val, err := d.r.Evaluate(d.model)
		if err != nil {
			t.Error(err)
		}
		if val != d.model[KSymbol{"out"}] {
			t.Errorf("case %d: unexpected value: %v", i, val)
		}
	}
}

func TestPop(t *testing.T) {
	p := KSymbol{"P"}
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
	p := KSymbol{"P"}
	q := KSymbol{"Q"}

	kb := KAnd{[]Prop{p, q}}
	val, err := ModelCheck(kb, q)
	if err != nil {
		t.Error(err)
	}
	if val != true {
		t.Errorf("unexpected value: %v", val)
	}
}

package bitset

import "testing"

func TestVect(t *testing.T) {
	b := New()
	b.Set(1)
	if b.Ones() != 1 {
		t.Error(b.String())
	}
	b.Set(6)
	if b.Ones() != 2 {
		t.Error(b.String())
	}
	b.Set(62)
	if !b.Has(62) {
		t.Error(b.String())
	}
	if b.Ones() != 3 {
		t.Error(b.String())
	}
	c := New()
	c.Set(62)
	if b.CountIntersect(c) != 1 {
		t.Errorf("\nIntesect count:%d\nb:%s\nc:%s\n", b.CountIntersect(c), b.String(), c.String())
	}
	c.Set(1)
	if b.CountIntersect(c) != 2 {
		t.Errorf("\nIntesect count:%d\nb:%s\nc:%s\n", b.CountIntersect(c), b.String(), c.String())
	}
	c.Clear(62)
	if b.CountIntersect(c) != 1 {
		t.Errorf("\nIntesect count:%d\nb:%s\nc:%s\n", b.CountIntersect(c), b.String(), c.String())
	}
}

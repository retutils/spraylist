//Package mutate. Provide utils to mutate passwords in list
package transform

import (
	"testing"
)

func TestTransforms(t *testing.T) {
	w := "Nowember"
	if Toggle(w) != "NoWeMbEr" {
		t.Error(Toggle(w))
	}
	if Lower(w) != "nowember" {
		t.Error(Lower(w))
	}
	if Uper(w) != "NOWEMBER" {
		t.Error(Uper(w))
	}
	if ToTitle("nowember") != "Nowember" {
		t.Error(ToTitle("nowember"))
	}
	if AppendSuffix(w, "!") != "Nowember!" {
		t.Error(AppendSuffix(w, "!"))
	}
	if AppendSuffix("Nowember1", "2011") != "Nowember1" {
		t.Error(AppendSuffix("Nowember1", "2011"))
	}
	if Replace(w, "o", "0") != "N0wember" {
		Replace("N0wember", "0", "0")
	}
	if Prepend("Password", "#") != "#Password" {
		Prepend("Password", "#")
	}
	if Prepend("12Password", "1") != "Password" {
		Prepend("12Password", "1")
	}
	if RussToEngKey("Сентябрь2020") != "Ctynz,hm2020" {
		t.Error(RussToEngKey("Сентябрь2020"))
	}

}

func TestUnuque(t *testing.T) {
	strSl := []string{"1", "1", "1", "3", "5", "7", "9", "10", "10", "1", "1", "1", "1", "1", "1", "1", "1",
		"19", "8", "9", "90", "0", "0", "0", "0", "0", "0", "10", "10"}
	res := UniqueStrings(strSl)
	if len(res) != 10 {
		t.Errorf("incorrect result: %v", res)
	}

}

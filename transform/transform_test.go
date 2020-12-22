//Package mutate. Provide utils to mutate passwords in list
package tranform

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
	if AppendSuffix(w, "!") != "Nowember1" {
		t.Error(AppendSuffix(w, "!"))
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

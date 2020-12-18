package policy

import "testing"

func TestMatrch(t *testing.T) {
	p, err := New(3, 9, "lower", "upper", "number", "nonAlpha")
	if err != nil {
		t.Fatal(err)
	}
	password := "password"
	if p.Match(password) {
		t.Error("Must not match policy")
	}
	password = "Password1!"
	if !p.Match(password) {
		t.Error("Must match policy")
	}
	password = "Password1"
	if !p.Match(password) {
		t.Error("Must mutch policy")
	}
	password = "Pard1"
	if p.Match(password) {
		t.Error("Must not match policy, by length")
	}
}

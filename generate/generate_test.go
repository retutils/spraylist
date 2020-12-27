package generate

import "testing"

func TestBasic(t *testing.T) {
	dict, err := New("../cfg/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	dict.addMonth("January", "2020", "20")
	if _, ok := dict.Result["January"]; !ok {
		t.Error("monnth not present")
	}
	if _, ok := dict.Result["January2020"]; !ok {
		t.Error("monnth year not present")
	}
	if _, ok := dict.Result["January@2020"]; !ok {
		t.Error("monnth year with delimeter not present")
	}
	if _, ok := dict.Result["January20"]; !ok {
		t.Error("monnth year short not present")
	}
	if _, ok := dict.Result["January@20"]; !ok {
		t.Error("monnth year short with delimeter not present")
	}
}

func TestDict(t *testing.T) {
	dict, err := New("../cfg/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	dict.MakeDictonary()
	if _, ok := dict.Result["November@20"]; !ok {
		t.Error("monnth year short with delimeter not present")
	}
	if _, ok := dict.Result["Ctynz,hm@2020"]; !ok {
		t.Error("monnth year russian with eng keyboard not present")
	}
	if _, ok := dict.Result["Password2020"]; !ok {
		t.Error("Common word with year not present")
	}
	dict.addReplaces("Password2020")
	if _, ok := dict.Result["Pa$$word2020"]; !ok {
		t.Error("Replaced common word  with year not present")
	}
	dict.addSuffixes("Pa$$word2020")
	if _, ok := dict.Result["Pa$$word2020!"]; !ok {
		t.Error("Suffix not present")
	}
	dict.addWithTransforms("Password2020")
	if _, ok := dict.Result["WiNtEr2020!"]; !ok {
		t.Error("Toogle transform not present")
	}
	if _, ok := dict.Result["winter2020!"]; !ok {
		t.Error("Lover transform not present")
	}
	dict.Write()
}

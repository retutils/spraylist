package cfg

import (
	"testing"
)

func TestPolicy(t *testing.T) {
	c, err := New("config.toml")
	if err != nil {
		t.Fatal(err)
	}
	if c.Policy.CharacterGroup[0] != "lower" {
		t.Error("Wrong char group", c.Policy.CharacterGroup[0])
	}
	if c.Policy.Lehgth != 9 {
		t.Error("Wrong pass length", c.Policy.Lehgth)
	}
	if c.Policy.RequiredChargroups != 3 {
		t.Error("Wrong Required Char groups", c.Policy.RequiredChargroups)
	}
	if c.Policy.Expire != 90 {
		t.Error("Wrong Expiration", c.Policy.Expire)
	}

}

func TestRules(t *testing.T) {
	c, err := New("config.toml")
	if err != nil {
		t.Fatal(err)
	}
	if c.Rules.Delimeters[0] != "@" {
		t.Error(c.Rules.Delimeters[0])
	}
	if c.Rules.Suffixes[0] != "1" {
		t.Error(c.Rules.Suffixes[0])
	}
	if c.Rules.Transforms[0] != "Lower" {
		t.Error(c.Rules.Transforms[0])
	}
	if c.Replace["s"] != "$" {
		t.Error(c.Replace["s"])
	}

}

func TestMounts(t *testing.T) {
	c, err := New("config.toml")
	if err != nil {
		t.Fatal(err)
	}
	if c.Month["January"][0] != "January" {
		t.Error("Wrong month", c.Month["January"])
	}
	if c.MonthRus["January"][0] != "Январь" {
		t.Error("Wrong month", c.MonthRus["January"])
	}
	m := c.MonthsYear()
	if len(m) < 4 {
		t.Error(m)
	}
}

func TestCommon(t *testing.T) {
	c, err := New("config.toml")
	if err != nil {
		t.Fatal(err)
	}
	if c.Company.Names[0] != "EPAM" {
		t.Error("Wrong company name", c.Company.Names[0])
	}
	if c.Common.Words[0] != "Password" {
		t.Error(c.Common.Words[0])
	}
}

//Package policy implements simple password policy matcher
//could have been replaced by a simple regex ;)
package policy

import (
	"fmt"
	"strings"

	"github.com/spraylist/bitset"
)

const (
	lower = iota
	upper
	number
	nonAlpha
)

//Policy defenetion
type Policy struct {
	length     int
	characters *bitset.Bitset
	groups     *bitset.Bitset
	groupsMin  int
}

//Match given policy
func (p *Policy) Match(password string) bool {
	p.characters = bitset.New()
	for _, value := range password {
		switch {
		case value >= 'a' && value <= 'z':
			p.characters.Set(lower)
		case value >= 'A' && value <= 'Z':
			p.characters.Set(upper)
		case value >= '0' && value <= '9':
			p.characters.Set(number)
		default:
			p.characters.Set(nonAlpha)
		}
	}
	return p.groups.CountIntersect(p.characters) >= p.groupsMin && len(password) >= p.length
}

//New (3,9,lower,upper,numeric,nonAlpha)
//charactersets: lower, upper, number , nonAlpha
//length : miminal password length
//required: minimum char group reqirements
//Example: passwords for Domain accounts (including temporary passwords) should be at least 9 alphanumeric characters.
//The password should fulfill at least three of the following for character group:
//English uppercase characters (A through Z);
// English lowercase characters (a through z);
//Numerals: (0 through 9);
//Non-alphabetic characters: (such as !, $, #,%).
func New(required, length int, charactersets ...string) (p *Policy, err error) {
	p = &Policy{length: length, groups: bitset.New(), groupsMin: required}
	for _, v := range charactersets {
		switch {
		case strings.ToLower(v) == "lower":
			p.groups.Set(lower)
		case strings.ToLower(v) == "upper":
			p.groups.Set(upper)
		case strings.ToLower(v) == "number":
			p.groups.Set(number)
		case strings.ToLower(v) == "nonalpha":
			p.groups.Set(nonAlpha)
		default:
			return nil, fmt.Errorf("invalid argument: %s", v)
		}

	}
	return p, nil

}

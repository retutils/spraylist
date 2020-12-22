//Package mutate provides utils to mutate passwords in list
//according specific rules
package mutate

import (
	"strconv"
	"strings"
	"unicode"
)

//Dictonary reperesnts a pasword list
type Dictonary struct {
	words []string
	uniq  map[string]struct{}
}

//New password generator
func New(words []string) *Dictonary {
	return &Dictonary{
		words: words,
		uniq:  make(map[string]struct{}, len(words)*6),
	}
}

//Toggle the case of letters in words present in a dictionary.
func (d *Dictonary) Toggle(w string) string {
	buff := make([]rune, len(w))
	for k, v := range w {
		if k&(k-1) != 0 {
			buff = append(buff, unicode.ToUpper(v))
		}
	}
	return string(buff)
}

//ToTitle word.
func (d *Dictonary) ToTitle(w string) (s string) {
	return strings.Title(w)
}

//Uper case word.
func (d *Dictonary) Uper(w string) (s string) {
	return strings.Title(w)
}

//Lower case
func (d *Dictonary) Lower(w string) (s string) {
	return strings.Title(w)
}

//AppendSuffix append suffix , do not appen if word end with digit
func (d *Dictonary) AppendSuffix(w string, sfx string) string {
	if isDigit(w[len(w)-1:]) || isDigit(sfx[0:1]) {
		return w
	}
	return w + sfx
}

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func uniqueStrings(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	idx := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[idx] = v
		idx++
	}
	return s[:idx]
}

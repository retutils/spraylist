//Package tranform provides utils to tranform passwords in list
//according specific rules
package tranform

import (
	"strconv"
	"strings"
	"unicode"
)

//Toggle the case of letters in words present in a dictionary.
func Toggle(w string) (s string) {
	var b strings.Builder
	b.Grow(len(w))
	even := true
	for _, v := range w {
		if even {
			b.WriteRune(unicode.ToUpper(v))
			even = false
		} else {
			b.WriteRune(v)
			even = true
		}
	}
	return b.String()
}

//ToTitle word.
func ToTitle(w string) (s string) {
	return strings.Title(w)
}

//Uper case word.
func Uper(w string) (s string) {
	return strings.ToUpper(w)
}

//Lower case
func Lower(w string) (s string) {
	return strings.ToLower(w)
}

//AppendSuffix apends suffix, do not append digit suffix if word end with digit
func AppendSuffix(w string, sfx string) string {
	if IsDigit(w[len(w)-1:]) && IsDigit(sfx[0:1]) {
		return w
	}
	var b strings.Builder
	b.Grow(len(w) + len(sfx))
	b.WriteString(w)
	b.WriteString(sfx)
	return b.String()
}

//Prepend word with preffix, do not append digit prefix if word start with digit
func Prepend(w string, pfx string) string {
	if IsDigit(w[0:1]) && IsDigit(pfx[len(pfx)-1:]) {
		return w
	}
	var b strings.Builder
	b.Grow(len(w) + len(pfx))
	b.WriteString(pfx)
	b.WriteString(w)
	return b.String()
}

//Replace single char
func Replace(w string, c, r string) string {
	var b strings.Builder
	b.Grow(len(w))
	for _, v := range w {
		if v == rune(c[0]) {
			b.WriteString(r)
		} else {
			b.WriteRune(v)
		}
	}
	return b.String()
}

//IsDigit retutns true if string is integer
func IsDigit(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

//UniqueStrings return slices with ubique strings
func UniqueStrings(s []string) []string {
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

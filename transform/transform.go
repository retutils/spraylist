//Package tranform provides utils to tranform passwords in list
//according specific rules
package tranform

import (
	"strconv"
	"strings"
	"unicode"
)

//Toggle the case of letters in words present in a dictionary.
func Toggle(w string) string {
	buff := make([]rune, len(w))
	even := true
	for _, v := range w {
		if even {
			buff = append(buff, unicode.ToUpper(v))
			even = false
		} else {
			buff = append(buff, v)
			even = true
		}
	}
	return string(buff)
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
	return strings.Title(w)
}

//AppendSuffix apends suffix, do not append digit suffix if word end with digit
func AppendSuffix(w string, sfx string) string {
	if IsDigit(w[len(w)-1:]) || IsDigit(sfx[0:1]) {
		return w
	}
	return w + sfx
}

//Prepend word with preffix, do not append digit prefix if word start with digit
func Prepend(w string, sfx string) string {
	if IsDigit(w[len(w)-1:]) || IsDigit(sfx[0:1]) {
		return w
	}
	return w + sfx
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

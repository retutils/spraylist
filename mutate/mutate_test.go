//Package mutate. Provide utils to mutate passwords in list
package mutate

import (
	"testing"
)

func TestUnui(t *testing.T) {
	strSl := []string{"1", "1", "1", "3", "5", "7", "9", "10", "10", "1", "1", "1", "1", "1", "1", "1", "1",
		"19", "8", "9", "90", "0", "0", "0", "0", "0", "0", "10", "10"}
	res := uniqueStrings(strSl)
	if len(res) != 10 {
		t.Errorf("incorrect result: %v", res)
	}

}

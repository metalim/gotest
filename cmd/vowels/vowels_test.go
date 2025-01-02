package vowels

import (
	"fmt"
	"testing"
)

func TestVowelStrings(t *testing.T) {
	words := []string{"apple", "orange", "banana", "pear"}
	queries := [][]int{{0, 2}, {1, 3}}
	ans := vowelStrings(words, queries)
	fmt.Println(ans)
}

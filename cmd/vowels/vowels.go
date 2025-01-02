package vowels

import (
	"strings"
)

const vowels = "aeiou"

func vowelStrings(words []string, queries [][]int) []int {
	counts := make([]int, len(words))
	var count int
	for i, w := range words {
		if strings.ContainsRune(vowels, rune(w[0])) && strings.ContainsRune(vowels, rune(w[len(w)-1])) {
			count++
		}
		counts[i] = count
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		var left int
		if q[0] > 0 {
			left = counts[q[0]-1]
		}
		right := counts[q[1]]
		ans[i] = right - left
	}
	return ans
}

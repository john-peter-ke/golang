package leetcode

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	result := []string{}

	for i := len(words) - 1; i >= 0; i-- {
		result = append(result, words[i])
	}

	return strings.Join(result, " ")
}

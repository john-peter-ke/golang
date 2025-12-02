package leetcode

import "strings"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var longest int = 1
	runes := []rune(s)
	n := len(runes)

	for j := 0; j <= n; j++ {
		substring := makeDistinctStr(string(runes[j:]))
		if len(substring) > longest {
			longest = len(substring)
		}

	}

	return longest
}

func makeDistinctStr(s string) string {
	distinct := make(map[rune]bool)
	var distinctChar strings.Builder
	for _, char := range s {
		_, found := distinct[char]
		if !found {
			distinct[char] = true
			distinctChar.WriteString(string(char))
		}
	}

	return distinctChar.String()
}

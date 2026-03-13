package leetcode

import "strings"

func isPalindrome(s string) bool {
	trimed := strings.ReplaceAll(s, " ", "")
	b := []byte(trimed)
	start, end := 0, len(b)-1
	for end > start {
		first := string(b[start])
		last := string(b[end])
		if !strings.EqualFold(strings.ToLower(first), strings.ToLower(last)) {
			return false
		}

		start++
		end--
	}

	return true
}

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

func lengthOfLongestSubstringWindow(s string) int {
	set := make(map[byte]bool)
	maxLen, left := 0, 0

	for right := 0; right < len(s); right++ {
		for set[s[right]] {
			delete(set, s[left])
			left++
		}
		set[s[right]] = true
		if (right - left + 1) > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func lengthOfLongestSubstringOptimized(s string) int {
	lastSeen := make(map[byte]int)
	maxLen, left := 0, 0

	for right := 0; right < len(s); right++ {
		// If char exists and is within current window, jump left
		if index, found := lastSeen[s[right]]; found && index >= left {
			left = index + 1
		}

		lastSeen[s[right]] = right

		if (right - left + 1) > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

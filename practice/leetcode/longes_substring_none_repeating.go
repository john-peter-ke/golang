package leetcode

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

func lengthOfLongestSubstringHashTable(s string) int {
	maxLen := 0

	for i := 0; i < len(s); i++ {
		set := make(map[byte]byte)
		substrLen := 0

		for j := i; j < len(s); j++ {
			if _, ok := set[s[j]]; ok {
				break
			}
			substrLen = (j - i) + 1
			set[s[j]] = s[j]
		}

		maxLen = max(maxLen, substrLen)

	}

	return maxLen
}

package leetcode

func lengthOfLongestSubstring(s string) int {
	// charIndexMap stores the last seen position of each character
	charIndexMap := make(map[byte]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		// If we've seen the character and it's within our current window
		if index, found := charIndexMap[char]; found && index >= left {
			// Jump the left pointer to the right of the previous occurrence
			left = index + 1
		}

		// Update the character's last seen index
		charIndexMap[char] = right

		// Calculate current window size and update maxLen
		currentWindowSize := right - left + 1
		if currentWindowSize > maxLen {
			maxLen = currentWindowSize
		}
	}

	return maxLen
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

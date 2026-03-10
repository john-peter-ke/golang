package leetcode

func longestPalindromeSubstring(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1

	for i := 0; i < len(s); i++ {
		// Case 1: Odd length palindromes (e.g., "aba" center at 'b')
		l1, r1 := expand(s, i, i)
		if r1-l1+1 > maxLength {
			start, maxLength = l1, r1-l1+1
		}

		// Case 2: Even length palindromes (e.g., "abba" center between 'b' and 'b')
		l2, r2 := expand(s, i, i+1)
		if r2-l2+1 > maxLength {
			start, maxLength = l2, r2-l2+1
		}
	}

	return s[start : start+maxLength]
}

func expand(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// Return indices of the last valid palindrome found before the loop failed
	return l + 1, r - 1
}

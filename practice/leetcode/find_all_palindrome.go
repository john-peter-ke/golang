package leetcode

func findAllPalindrome(s string) map[string]bool {
	palindromes := make(map[string]bool)
	if len(s) == 0 {
		return palindromes
	}

	for i := 0; i < len(s); i++ {
		expandCenter(s, i, i, palindromes)
		expandCenter(s, i, i+1, palindromes)
	}

	return palindromes
}

func expandCenter(s string, left, right int, palindromes map[string]bool) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		palindrome := s[left : right+1]
		palindromes[palindrome] = true
		left--
		right--
	}
}

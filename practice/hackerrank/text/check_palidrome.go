package text

import "strings"

func isAlphabeticPalindrome(code string) bool {
	normal := ""
	reversed := ""
	for _, r := range code {
		if isASCIILetter(r) {
			normal = normal + string(r)
			reversed = string(r) + reversed
		}
	}

	return strings.EqualFold(normal, reversed)
}

func isASCIILetter(c rune) bool {
	// Check for uppercase letters
	if c >= 'A' && c <= 'Z' {
		return true
	}
	// Check for lowercase letters
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}

package text

import "strings"

func maxDistinctSubstringLengthInSessions(sessionString string) int32 {
	if len(sessionString) == 0 || len(sessionString) == 1 {
		return 0
	}

	unique := make(map[string]int)
	for _, val := range sessionString {

		str := strings.ToLower(string(val))
		if found, ok := unique[str]; ok && found > 0 {
			//skip
			continue
		}

		unique[str] += 1

	}

	return int32(len(unique))
}

func isLetter(c rune) bool {
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

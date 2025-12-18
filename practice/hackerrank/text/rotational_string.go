package text

import "strings"

func isNonTrivialRotation(s1 string, s2 string) bool {
	// 1. Check for length equality
	if len(s1) != len(s2) {
		return false
	}

	// 2. Handle identical strings (if we want *non-identical* rotations)
	if s1 == s2 {
		return false // Exclude identical strings
	}

	// 3. Concatenate and search
	s1Double := s1 + s1
	return strings.Contains(s1Double, s2)
}

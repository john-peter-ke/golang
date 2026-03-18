package leetcode

func groupAnagrams(strs []string) [][]string {
	// map key: [26]int represents counts of 'a' through 'z'
	// map value: []string contains the actual anagrams
	anagramsMap := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, char := range s {
			// Subtracting 'a' gives us the index 0-25
			count[char-'a']++
		}
		// Since arrays are comparable in Go, we use them as keys directly
		anagramsMap[count] = append(anagramsMap[count], s)
	}

	// Convert map values to the required return format
	result := make([][]string, 0, len(anagramsMap))
	for _, group := range anagramsMap {
		result = append(result, group)
	}

	return result
}

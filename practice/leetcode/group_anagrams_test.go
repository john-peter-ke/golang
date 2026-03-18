package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	// Table-driven test cases
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "standard case",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
		},
		{
			name:     "empty string",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "single character",
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.input)

			// Because map iteration order is random, we sort the results
			// to compare 'got' and 'expected' accurately.
			if !compareAnagramGroups(got, tt.expected) {
				t.Errorf("groupAnagrams(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// Helper to normalize and compare slices of slices for testing
func compareAnagramGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	sortGroups := func(groups [][]string) {
		for _, g := range groups {
			sort.Strings(g)
		}
		sort.Slice(groups, func(i, j int) bool {
			return groups[i][0] < groups[j][0]
		})
	}
	sortGroups(a)
	sortGroups(b)
	return reflect.DeepEqual(a, b)
}

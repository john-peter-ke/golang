package aligorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SortTestCase struct {
	Name           string
	Input          []int
	ExpectedResult []int
	ExpectedError  error
}

type SearchTestCase struct {
	Name           string
	Input          []int
	Target         int
	ExpectedResult int
	ExpectedError  error
}

var sortCases = []SortTestCase{
	{
		Name:           "Small array",
		Input:          []int{1, 4, 5, 7, 2, 9, 3, 6, 8},
		ExpectedResult: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		ExpectedError:  nil,
	},
	{
		Name:           "Almost sorted array",
		Input:          []int{7, 3, 9, 12, 11},
		ExpectedResult: []int{3, 7, 9, 11, 12},
		ExpectedError:  nil,
	},
	{
		Name:           "Completely unsorted array",
		Input:          []int{4, 6, 2, 9, 3, 1, 5, 7, 8},
		ExpectedResult: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		ExpectedError:  nil,
	},
}

var searchCases = []SearchTestCase{
	{
		Name:           "Search 7",
		Input:          []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Target:         7,
		ExpectedResult: 6,
		ExpectedError:  nil,
	},
	{
		Name:           "Search 4",
		Input:          []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Target:         4,
		ExpectedResult: 3,
		ExpectedError:  nil,
	},
}

func ExecutesSortTestCases(t *testing.T, exefunc func(input []int) []int) {
	for i, tc := range sortCases {
		tcn := fmt.Sprintf("%d-%s", i, tc.Name)
		t.Run(tcn, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedResult, exefunc(tc.Input))
		})
	}
}

func ExecutesSearchTestCases(t *testing.T, exefunc func(input []int, searchItem int) (int, error)) {
	for i, tc := range searchCases {
		tcn := fmt.Sprintf("%d-%s", i, tc.Name)
		t.Run(tcn, func(t *testing.T) {
			actual, err := exefunc(tc.Input, tc.Target)
			assert.Equal(t, tc.ExpectedResult, actual)
			assert.Equal(t, tc.ExpectedError, err)
		})
	}
}

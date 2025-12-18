package text

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func Staircase(n int32) {
	size := int(n)
	for i := 1; i <= size; i++ {
		fmt.Print(spaces(size-i), symbol(i), "\n")
	}
}

func spaces(n int) string {
	var str strings.Builder
	for i := 1; i <= n; i++ {
		str.WriteString(" ")
	}

	return str.String()
}

func symbol(n int) string {
	var str strings.Builder
	for i := 1; i <= n; i++ {
		str.WriteString("#")
	}

	return str.String()
}

package main

import (
	"bufio"
	"fmt"
	"io"
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

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

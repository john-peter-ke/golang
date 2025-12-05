package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	printStaircase()
}

func printStaircase() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	Staircase(n)
}

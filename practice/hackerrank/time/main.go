package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	timeConversion()
}

func timeConversion() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := TimeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	print("how many times should I say 'Hello world!': ")
	// stop listening and return at \n (newline)
	inpByte, _, _ := reader.ReadLine()
	inp := string(inpByte)

	// get rid of the trailing \n
	inp = strings.Replace(inp, "\n", "", -1)

	times, _ := strconv.Atoi(inp)
	for i := 0; i < times; i++ {
		println("Hello world!")
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	countLines := flag.Bool("l", false, "Count lines instead of words")
	flag.Parse()

	fmt.Println(count(os.Stdin, *countLines))
}

func count(r io.Reader, shouldCountLines bool) int {
	scanner := bufio.NewScanner(r)

	if !shouldCountLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
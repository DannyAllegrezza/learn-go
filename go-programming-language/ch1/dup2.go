// Dup2 prints the count and text of lines that appear more than once in the input
// It reads from `stdin` OR from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	if filePathsPresent() {
		// read through the files, and count their lines
		files := os.Args[1:]

		for _, arg := range files {
			file, err := os.Open(arg)

			if err != nil {
				// uh oh, couldn't read this file. try to move on to the next file provided
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	} else {
		// no files provided, read stdin
		countLines(os.Stdin, counts)
	}

	printCounts(counts)
}

func filePathsPresent() bool {
	return len(os.Args[1:]) > 0
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}
}

func printCounts(counts map[string]int) {
	// now that we're done, display results to the user
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}

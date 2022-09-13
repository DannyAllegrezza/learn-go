// Dup2 prints the count and text of lines that appear more than once in the input
// It reads from `stdin` OR from a list of named files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No filename provided")
		return
	}

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile returns a byte slice that must be converted into a string
		// so it can be split by strings.Split
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprint(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

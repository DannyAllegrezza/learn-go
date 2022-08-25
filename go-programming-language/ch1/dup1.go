// Dup1 prints the text of each line that appears more than
// once in the standard input stream preceded by its count
//
// go run dup1
// hey
// hey
// hi
// <terminate with ctrl+d
// 2     hey
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// loop through all input, line by line and place
	// the value of the line as the key in the map and increment its value
	for input.Scan() {
		counts[input.Text()]++
	}

	// now that we're done, display results to the user
	for line, n := range counts {
		if n > 1 {
			//fmt.Printf("%d\t%s\n", n, line)
			fmt.Println(n, line)
		}
	}
}

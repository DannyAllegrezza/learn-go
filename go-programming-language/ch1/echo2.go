// Echo2 prints its command-line arguments
// ex: go run echo1.go hello world, friends
// will print "hello world, friends"
package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, separator string
	s, separator := "", ""

	// os.Args[1:] is interesting in Go
	for _, arg := range os.Args[1:] {
		s += separator + arg
		separator += " "
	}

	fmt.Println(s)
}

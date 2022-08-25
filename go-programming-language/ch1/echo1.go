// Echo1 prints its command-line arguments
// ex: go run echo1.go hello world, friends
// will print "hello world, friends"
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, separator string

	for index := 1; index < len(os.Args); index++ {
		s += separator + os.Args[index]
		separator = " "
	}

	fmt.Println(s)
}

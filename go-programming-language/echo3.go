// Echo3 prints its command-line arguments
// ex: go run echo3.go hello world, friends
// will print "hello world, friends"
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println("-------")

	fmt.Println(os.Args[1:])

	fmt.Println(os.Args)

	fmt.Println("-------")

	fmt.Println(len(os.Args), "len(os.Args)")
}

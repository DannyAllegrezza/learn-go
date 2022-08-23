// Echo3 prints its command-line arguments
// ex: go run echo3.go hello world, friends
// will print "hello world, friends"
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	fmt.Println("-------")

	fmt.Println("os.Args[1:] --->", os.Args[1:])

	fmt.Println("-------")

	fmt.Println("len(os.Args) --->", len(os.Args))

	fmt.Println("Exercise 1.1")
	fmt.Println(os.Args[0:])

	fmt.Println("\nExercise 1.2 - Print index and value of each of its arguments, one per line")
	for index, element := range os.Args {
		fmt.Println(index, element)
	}
}

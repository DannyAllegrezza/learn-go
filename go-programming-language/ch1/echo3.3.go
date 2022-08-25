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

	start = time.Now()
	potentiallyInefficientParsing()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func potentiallyInefficientParsing() {
	var s, separator string

	for index := 1; index < len(os.Args); index++ {
		s += separator + os.Args[index]
		separator = " "
	}

	fmt.Println(s)
}

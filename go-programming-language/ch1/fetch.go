// Fetch prints the content found at the provided URL received via stdin
//
// go run fetch.go https://www.dannyallegrezza.com
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		fetchUrl(url)
	}
}

func fetchUrl(url string) {
	if !hasHttpPrefix(url) {
		url = "http://" + url
	}

	fmt.Println("Fetching URL ", url)
	response, err := http.Get(url)

	if err != nil {
		log.Fatalln("fetch failure: ", err)
	}

	// exercise 1.6 - read response into a buffer
	// body, err := ioutil.ReadAll(response.Body)
	// response.Body.Close()
	//
	// if err != nil {
	// 	fmt.Println("fetch: reading", url, err)
	// }
	// fmt.Printf("%s", body)

	// exercise 1.7 - use io.Copy() instead
	if _, err := io.Copy(os.Stdout, response.Body); err != nil {
		log.Fatal("fetch: reading", url, err)
	}
}

func hasHttpPrefix(url string) bool {
	return strings.HasPrefix(url, "http://")
}

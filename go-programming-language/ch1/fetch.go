// Fetch prints the content found at the provided URL received via stdin
//
// go run fetch.go https://www.dannyallegrezza.com
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		fetchUrl(url)
	}
}

func fetchUrl(url string) {
	response, err := http.Get(url)

	if err != nil {
		log.Fatalln("fetch failure: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		fmt.Println("fetch: reading", url, err)
	}
	fmt.Printf("%s", body)
}

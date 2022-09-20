// fetchall fetches all provided URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	startingTime := time.Now()
	channel := make(chan string)

	// for each provided URL (via args), spin up a separate go routine
	// we pass along a single channel of strings as a way for our separate
	// goroutines to communicate
	for _, url := range os.Args[1:] {
		go fetch(url, channel)
	}

	for range os.Args[1:] {
		fmt.Println(<-channel) // receive from channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(startingTime).Seconds())
}

// fetch performs a HTTP GET request of the provided `url` param and
// returns the total byte size of the HTML as well as the elapsed time in seconds
// that it took to perform the request to the provided channel `ch`
func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nBytes, err := io.Copy(io.Discard, response.Body)
	response.Body.Close() // we don't want to leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nBytes, url)
}

// Server1 is a minimal echo web server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handleCount)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	count++
	mutex.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// handleCount echoes the number of calls that the server has received so far
func handleCount(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Println("Count", count)
	mutex.Unlock()
}

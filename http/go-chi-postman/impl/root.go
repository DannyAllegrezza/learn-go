package impl

import (
	"net/http"
)

type RootImpl struct{}

// Get /
func (root *RootImpl) Get(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
	w.Header().Set("Content-Type", "text/plain")
	helloWorld := []byte("hello world")

	w.Write(helloWorld)
	w.WriteHeader(200)
}

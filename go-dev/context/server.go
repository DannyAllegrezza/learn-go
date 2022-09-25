package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"example.com/context/github"
	"example.com/context/userip"
)

func main() {
	http.HandleFunc("/stars", handleGithubSearch)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleGithubSearch(w http.ResponseWriter, req *http.Request) {

	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// the HTTP request includes a timeout, so create a context
		// that is canceled automatically when the timeout expires
		fmt.Printf("Request has timeout value")
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	query := req.FormValue("username")
	if query == "" {
		http.Error(w, "No username provided", http.StatusBadRequest)
		return
	}

	// Store the user IP in ctx for use by code in other packages.
	userIP, err := userip.FromRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)

	start := time.Now()
	results, err := github.GetUsersStarredResults(ctx, query)
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := resultsTemplate.Execute(w, struct {
		Results          github.UserStarredResults
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}

}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Name}} - <a href="{{.HTMLURL}}">{{.HTMLURL}}</a></li>
  {{end}}
  </ol>
</body>
</html>
`))

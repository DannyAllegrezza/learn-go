package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"strconv"
)

// Album represents data about a musical record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Database is the interface used by the server to load and store albums.
type Database interface {
	GetAlbums() ([]Album, error)
	GetAlbumById(id string) (Album, error)
	AddAlbum(album Album) error
}

type Server struct {
	db     Database
	log    *log.Logger
	router *http.ServeMux
}

func main() {
	// Allow the user to optionally specify the port to run the web server on (ex: ./web-service-2 -port 1222)
	var port int
	flag.IntVar(&port, "port", 1337, "port to listen on")
	flag.Parse()

	// Create our in-memory DB with a few test albums
	db := NewInMemoryDatabase()

	server := NewServer(db, log.Default())
	log.Printf("listening on http://localhost:%d", port)
	http.ListenAndServe(":"+strconv.Itoa(port), server)
}

// Handlers
// These hang right off the Server, giving us easy access to the servers dependencies
func (s *Server) handleListAlbums() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		a, _ := s.db.GetAlbums()

		s.writeJSON(w, 200, a)
	}
}

// Server functions

// writeJSON marshals v to JSON and writes it to the response, handling
// errors as appropriate. It also sets the Content-Type header to
// "application/json".
func (s *Server) writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		s.log.Printf("error marshaling JSON: %v", err)
		http.Error(w, `{"error":"`+ErrorInternal+`"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		// Very unlikely to happen, but log any error (not much more we can do)
		s.log.Printf("error writing JSON: %v", err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(db Database, log *log.Logger) *Server {
	server := &Server{
		db:     db,
		log:    log,
		router: http.NewServeMux(),
	}
	server.routes()
	return server
}

// InMemoryDatabase is an implementation of our Database interface that uses a simple
// in-memory slice of Albums.
type InMemoryDatabase struct {
	albums []Album
}

// Database functions

func (db *InMemoryDatabase) GetAlbums() ([]Album, error) {
	return db.albums, nil
}

func (db *InMemoryDatabase) GetAlbumById(id string) (Album, error) {
	for _, v := range db.albums {
		if v.ID == id {
			return v, nil
		}
	}

	return Album{}, ErrDoesNotExist
}

func (db *InMemoryDatabase) AddAlbum(album Album) error {
	// TODO: Do we have this in our slice already?
	db.albums = append(db.albums, album)

	return nil
}

func NewInMemoryDatabase() *InMemoryDatabase {
	var albums = []Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return &InMemoryDatabase{
		albums: albums,
	}
}

var (
	ErrDoesNotExist  = errors.New("does not exist")
	ErrAlreadyExists = errors.New("already exists")
)

const (
	ErrorAlreadyExists    = "already-exists"
	ErrorDatabase         = "database"
	ErrorInternal         = "internal"
	ErrorMalformedJSON    = "malformed-json"
	ErrorMethodNotAllowed = "method-not-allowed"
	ErrorNotFound         = "not-found"
	ErrorValidation       = "validation"
)

package main

func (s *Server) routes() {
	s.router.HandleFunc("/albums", s.handleListAlbums())
}

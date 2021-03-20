package main

import "net/http"

func (s *server) routes() {
	s.mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// s.mux.HandleFunc("/", s.handleIndex())
	s.mux.HandleFunc("/dashboard", s.handleDashboard())
	s.mux.HandleFunc("/profile/show", s.handleProfileShow())
	s.mux.HandleFunc("/profile/edit", s.handleProfileEdit())
}

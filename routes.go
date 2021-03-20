package main

func (s *server) routes() {
	s.mux.HandleFunc("/", s.handleIndex())
	s.mux.HandleFunc("/dashboard", s.handleDashboard())
	s.mux.HandleFunc("/profile/show", s.handleProfileShow())
	s.mux.HandleFunc("/profile/edit", s.handleProfileEdit())
}

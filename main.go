package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/ChrisJohnson-83/web-frontend-demo/public"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func GetEnvDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func run(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		port = flags.String("port", GetEnvDefault("PORT", "8080"), "port to listen on")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	addr := fmt.Sprintf("0.0.0.0:%s", *port)
	// addr := fmt.Sprintf("localhost:%s", *port)
	srv, err := newServer()
	if err != nil {
		return err
	}
	fmt.Printf("cccs listening on :%s\n", *port)
	return http.ListenAndServe(addr, srv)
}

type server struct {
	mux *http.ServeMux
}

func newServer() (*server, error) {
	srv := &server{
		mux: http.NewServeMux(),
	}
	srv.routes()
	return srv, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) handleDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := public.Menus{
			public.Menu{
				Item: "Jack",
			},
			public.Menu{
				Item: "Jill",
			},
		}
		p := public.DashboardParams{
			Menus:   m,
			Title:   "Dashboard",
			Message: "Hello from dashboard",
		}
		public.Dashboard(w, p)
	}
}

func (s *server) handleProfileShow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := public.Menus{
			public.Menu{
				Item: "Jack",
			},
			public.Menu{
				Item: "Jill",
			},
		}
		p := public.ProfileShowParams{
			Menus:   m,
			Title:   "Profile Show",
			Message: "Hello Chris from your profile",
		}
		public.ProfileShow(w, p)
	}
}

func (s *server) handleProfileEdit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := public.Menus{
			public.Menu{
				Item: "Jack",
			},
			public.Menu{
				Item: "Jill",
			},
		}
		p := public.ProfileEditParams{
			Menus:   m,
			Title:   "Profile Edit",
			Message: "Hello from profile edit",
		}
		public.ProfileEdit(w, p)
	}
}

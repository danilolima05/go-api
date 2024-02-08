package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router        chi.Router
	PostHandlers  map[string]http.HandlerFunc
	GetHandlers   map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		PostHandlers:  make(map[string]http.HandlerFunc),
		GetHandlers:   make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddPostHandler(path string, handler http.HandlerFunc) {
	s.PostHandlers[path] = handler
}

func (s *WebServer) AddGetHandler(path string, handler http.HandlerFunc) {
	s.GetHandlers[path] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	for path, handler := range s.GetHandlers {
		s.Router.Get(path, handler)
	}

	for path, handler := range s.PostHandlers {
		s.Router.Post(path, handler)
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}

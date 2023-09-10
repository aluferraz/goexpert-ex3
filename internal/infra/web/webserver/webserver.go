package webserver

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RESTEndpoint struct {
	path string
	verb string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[RESTEndpoint]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[RESTEndpoint]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, verb string, handler http.HandlerFunc) {
	s.Handlers[RESTEndpoint{
		path: path,
		verb: verb,
	}] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() error {
	s.Router.Use(middleware.Logger)
	for restEndpointInfo, handler := range s.Handlers {
		path := restEndpointInfo.path
		switch verb := restEndpointInfo.verb; verb {
		case http.MethodGet:
			s.Router.Get(path, handler)
		case http.MethodPost:
			s.Router.Post(path, handler)
		case http.MethodPut:
			s.Router.Put(path, handler)
		case http.MethodPatch:
			s.Router.Patch(path, handler)
		case http.MethodDelete:
			s.Router.Delete(path, handler)
		default:
			return errors.New("invalid HTTP Verb")
		}

	}
	http.ListenAndServe(s.WebServerPort, s.Router)
	return nil
}

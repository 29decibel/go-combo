package gocombo

import (
	"fmt"
	"net/http"
	"regexp"
)

// Route is a struct for general Route infos
type Route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

// RegexpHandler is the handler struct for handle requests
type RegexpHandler struct {
	routes []*Route
}

// Handler is to handle requests
func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &Route{pattern, handler})
}

// HandleFunc is the func which handle requests
func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &Route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		fmt.Println("here is the given url:")
		fmt.Println(r.URL.Path)
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

package gocombo

import (
  "fmt"
  "net/http"
  "regexp"
)

type Route struct {
  pattern *regexp.Regexp
  handler http.Handler
}

type RegexpHandler struct {
  routes []*Route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
  h.routes = append(h.routes, &Route{pattern, handler})
}

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


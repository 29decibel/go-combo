package gocombo

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "regexp"
)


// it could be in the file system, it could be in the s3
const DirLocation = "~/projects/acxiom-js-library"

type Request struct {
  Resources []string
  Type string
}


func ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
  fmt.Println("do some crazy things here")
}


// retrieve contents of given resource
func (request Request)resourceContents(resourceName string) string {
  fileContents, _ := ioutil.ReadFile(resourceName)
  return string(fileContents)
}


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


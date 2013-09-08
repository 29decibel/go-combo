package main

import (
  "strings"
  "fmt"
  "log"
  "io"
  "net/http"
  "regexp"
  "github.com/29decibel/gocombo"
)

const Port = ":8123"

func EchoHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Path:")
  io.WriteString(w, req.URL.Path + "\n\n")
  io.WriteString(w, "URL:\n")
  io.WriteString(w, req.URL.RawQuery + "\n\n")
  resources := strings.Split(req.URL.RawQuery, "&")
  for _, resourceName := range resources {
    fmt.Println("got one ")
    io.WriteString(w, resourceName + "\n")
  }
}

func main() {
  fmt.Println("Start combo server here....")

  // create handlers here
  regHandler := gocombo.RegexpHandler{}
  re := regexp.MustCompile("/.*")

  regHandler.HandleFunc(re, EchoHandler)

  log.Fatal(http.ListenAndServe(Port, &regHandler))
}

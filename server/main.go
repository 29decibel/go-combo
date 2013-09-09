package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/29decibel/gocombo"
)

const Port = ":8123"

func main() {

  // TODO replace "./" into the given base path
  http.Handle("/", http.FileServer(http.Dir(gocombo.OptionValue("--base"))))

  http.HandleFunc("/combo", gocombo.ServeHTTP)

  // custom port
  port := ":" + gocombo.OptionValue("--port")
  if len(port) != 5 {
    port = Port
  }

  fmt.Println(fmt.Sprintf("Start YUI combo handler server(http://localhost%s) ...", port))
  log.Fatal(http.ListenAndServe(port, nil))
}

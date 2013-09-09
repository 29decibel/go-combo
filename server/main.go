package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/29decibel/gocombo"
)

const Port = ":8123"

func main() {
  fmt.Println("Start combo server here....")

  http.HandleFunc("/combo", gocombo.ServeHTTP)

  log.Fatal(http.ListenAndServe(Port, nil))
}

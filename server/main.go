package main

import (
  "fmt"
  "log"
  "flag"
  "net/http"
  "github.com/29decibel/gocombo"
)

func main() {

  var baseDirP *string = flag.String("base", "./yui3/build", "given the yui3 build file directory")
  var withVersionP *bool = flag.Bool("with-version", false, "if support version number")
  var portP *string = flag.String("port", ":8123", "service port, default will be 8123")
  flag.Parse()

  // set values of ComboConfig
  comboConfig := gocombo.ComboConfig{BaseDir: *baseDirP, WithVersion: *withVersionP, Port: *portP}
  gocombo.SetConfig(&comboConfig)


  // TODO replace "./" into the given base path
  http.Handle("/", http.FileServer(http.Dir(gocombo.ResourceDir())))

  http.HandleFunc("/combo", gocombo.ServeHTTP)

  fmt.Println(fmt.Sprintf("Start YUI combo handler server(http://localhost%s) ...", comboConfig.Port))
  log.Fatal(http.ListenAndServe(comboConfig.Port, nil))
}






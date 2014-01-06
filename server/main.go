package main

import (
	"flag"
	"fmt"
	"github.com/29decibel/gocombo"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	var baseDirP = flag.String("base", "./yui3/build", "given the yui3 build file directory")
	var withVersionP = flag.Bool("with-version", false, "if support version number")
	var portP = flag.String("port", ":8123", "service port, default will be 8123")
	flag.Parse()

	// set values of ComboConfig
	currentDir, _ := os.Getwd()
	comboConfig := gocombo.ComboConfig{BaseDir: filepath.Join(currentDir, *baseDirP), WithVersion: *withVersionP, Port: *portP}
	gocombo.SetConfig(&comboConfig)

	// TODO replace "./" into the given base path
	http.Handle("/", http.FileServer(http.Dir(comboConfig.BaseDir)))

	http.HandleFunc("/combo", gocombo.ServeHTTP)

	fmt.Println(fmt.Sprintf("Start YUI combo handler server(http://localhost%s) ...", comboConfig.Port))
	fmt.Println(fmt.Sprintf("Serving files from directory: %s", comboConfig.BaseDir))
	log.Fatal(http.ListenAndServe(comboConfig.Port, nil))
}

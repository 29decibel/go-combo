package gocombo

import (
  "fmt"
  "net/http"
  "strings"
  "io"
  "io/ioutil"
  "os"
)


// it could be in the file system, it could be in the s3
const DirLocation = "./default-yui-built-files"

type Request struct {
  Resources []string
  Type string
  BasePath string
}


func ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
  // get the resources
  resources := strings.Split(request.URL.RawQuery, "&")

  // create the combo request
  comboReq := Request{Resources: resources}

  // get the current dir
  for _, arg := range os.Args {
    fmt.Println(arg)
    options := strings.Split(arg, "=")
    if len(options) > 1 && options[0] == "--base" {
      comboReq.BasePath = options[1]
    }
  }

  responseWriter.Header().Set("Content-Type", "application/javascript; charset=utf-8")

  io.WriteString(responseWriter, comboReq.ResponseString())
}


func (request Request)ResponseString() string{
  contents := ""
  for _, resourceName := range request.Resources {
    contents += readFile(request, resourceName)
  }

  return contents
}

func readFile(request Request, resourceName string) string {
  parts := strings.Split(resourceName, "/")
  path := strings.Join(parts[1:], "/")

  var fileName string
  if len(request.BasePath) > 1 {
    fileName = request.BasePath + path
  } else {
    fileName = DirLocation + path
  }
  // get the file name
  contents, err := ioutil.ReadFile(fileName)
  if err != nil { panic(err) }
  return string(contents)
}



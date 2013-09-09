package gocombo

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
  "regexp"
)

// it could be in the file system, it could be in the s3
const (
	DirLocation    = "./default-yui-built-files"
	ContentType    = "Content-Type"
	JSContentType  = "application/javascript; charset=utf-8"
	CSSContentType = "text/css; charset=utf-8"
)

type Request struct {
	Resources []string
	Type      string
	BasePath  string
}

func ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// get the resources
	resources := strings.Split(request.URL.RawQuery, "&")

	// create the combo request
	comboReq := Request{Resources: resources}

	// get the current dir
  baseDir := OptionValue("--base")
  if len(baseDir) >0 {
		comboReq.BasePath = baseDir
  }

	// set header of the response
	acceptHeader := strings.Join(request.Header["Accept"], ",")
	if strings.Contains(acceptHeader, "css") {
		responseWriter.Header().Set(ContentType, CSSContentType)
	} else {
		responseWriter.Header().Set(ContentType, JSContentType)
	}

	io.WriteString(responseWriter, comboReq.ResponseString())
}


func updateImagePath(resourceName string, contents string) string {
  // replace url(./sorting.png) to
  // ....
  imageReg := regexp.MustCompile("url\\('?(.*?)'?\\)")

  contents = imageReg.ReplaceAllString(contents, "url(" + stripVersionNumber(resourceName) +"/$1)")
  return contents
}

func stripVersionNumber(name string) string {
	parts := strings.Split(name, "/")
	return strings.Join(parts[1:len(parts)-1], "/")
}

func (request Request) ResponseString() string {
	contents := ""

	// use a chanel to collect the resource contents
	contentsChanel := make(chan string)
	for _, resourceName := range request.Resources {
		go readFile(request, contentsChanel, resourceName)
	}

	// collect contents here
	for i := 0; i < len(request.Resources); i++ {
		contents += <-contentsChanel
	}

	return contents
}

// read file to chanel
func readFile(request Request, contentsChanel chan string, resourceName string) {
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
	if err != nil {
		panic(err)
	}
	contentsChanel <- updateImagePath(resourceName, string(contents))
}

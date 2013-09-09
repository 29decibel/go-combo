package gocombo

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// it could be in the file system, it could be in the s3
const (
	DirLocation   = "./default-yui-built-files"
	ContentType   = "Content-Type"
	JSContentType = "application/javascript; charset=utf-8"
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
	for _, arg := range os.Args {
		options := strings.Split(arg, "=")
		if len(options) > 1 && options[0] == "--base" {
			comboReq.BasePath = options[1]
		}
	}

	responseWriter.Header().Set(ContentType, JSContentType)

	io.WriteString(responseWriter, comboReq.ResponseString())
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
	contentsChanel <- string(contents)
}

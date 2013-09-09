package gocombo

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
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
	fmt.Printf("%s(%s): %s\n", request.Method, time.Now(), request.URL.Path)
	// get the resources
	resources := strings.Split(request.URL.RawQuery, "&")

	// create the combo request
	comboReq := Request{Resources: resources}

	// get the current dir
	baseDir := OptionValue("--base")
	if len(baseDir) > 0 {
		comboReq.BasePath = baseDir
	}

	// set header of the response
	acceptHeader := strings.Join(request.Header["Accept"], ",")
	if strings.Contains(acceptHeader, "css") {
		responseWriter.Header().Set(ContentType, CSSContentType)
	} else {
		responseWriter.Header().Set(ContentType, JSContentType)
	}

	// set cookie
	expires := time.Now().AddDate(3, 0, 0)
	responseWriter.Header().Set("Expires", expires.Format(time.UnixDate))

	io.WriteString(responseWriter, comboReq.ResponseString())
}

func updateImagePath(resourceName string, contents string) string {
	// replace url(./sorting.png) to
	// ....
	imageReg := regexp.MustCompile("url\\('?(.*?)'?\\)")

	_, path, _ := splitResourceName(resourceName)
	contents = imageReg.ReplaceAllString(contents, fmt.Sprintf("url(%s/$1)", path))
	return contents
}

// given resource name : 3.12.0/event-hover/event-hover-min.js
// return parts like [versionNumber][path][fileName]
// [3.12.0] [event-hovea] [event-hover-min.js]
func splitResourceName(name string) (string, string, string) {
	parts := strings.Split(name, "/")

	// pick the interesting part
	versionNumber := parts[0]
	path := strings.Join(parts[1:len(parts)-1], "/")
	fileName := parts[len(parts)-1]
	return versionNumber, path, fileName
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

package gocombo

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// it could be in the file system, it could be in the s3
const (
	ContentType    = "Content-Type"
	JSContentType  = "application/javascript; charset=utf-8"
	CSSContentType = "text/css; charset=utf-8"
)

var comboConfig ComboConfig

type ComboConfig struct {
	// servering files base directory
	BaseDir string
	// if with version number support
	WithVersion bool
	// server port
	Port string
}

// Request represent a combo resource request, either js or css
type ComboRequest struct {
	// names of resources
	Resources []string
	// the resource type
	Type string
}

// handler of combo request
// given http request, create a ComboRequest
func ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// get the resources
	resources := strings.Split(request.URL.RawQuery, "&")

	// create the combo request
	comboReq := ComboRequest{Resources: resources}

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

// set the combo config
func SetConfig(config *ComboConfig) {
	comboConfig = *config
}

// replace resource url in css with absolute paht
func updateImagePath(resourceName string, contents string) string {
	// replace url(./sorting.png) to
	// ....
	imageReg := regexp.MustCompile("url\\('?(.*?)'?\\)")

	version, path, _ := splitResourceName(resourceName)
	if comboConfig.WithVersion {
		path = fmt.Sprintf("%s/build/%s", version, path)
	}
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

// get the response string of all resources
func (request ComboRequest) ResponseString() string {
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
func readFile(request ComboRequest, contentsChanel chan string, resourceName string) {
	parts := strings.Split(resourceName, "/")
	path := strings.Join(parts[1:], "/")

	versionPart := []string{parts[0], "build"}
	if comboConfig.WithVersion {
		path = fmt.Sprintf("%s/%s", strings.Join(versionPart, "/"), path)
	}

	fileName := filepath.Join(comboConfig.BaseDir, path)

	// get the file name
	// or we can use os.Stat to check if the file exist
	// if _, err := os.Stat(filename); err == nil
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("Can not read file \"%s\" ! Please check if given --base path is correct.", fileName))
	}
	contentsChanel <- updateImagePath(resourceName, string(contents))
}

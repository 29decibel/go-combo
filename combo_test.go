package gocombo

import "testing"

// test resource names parsing
func TestResourceNamesParse(t *testing.T) {

}

// test resources reading
func TestResourceNameSplit(t *testing.T) {
	// define two resource names
	resourceName1 := "0.0.1/some-module/aa.js"
	resourceName2 := "0.0.1/some-module/aa.js"

	// create the request
	request := Request{Resources: []string{resourceName1, resourceName2}}
	if request.ResponseString() == "" {
		t.Log("success get contents of resource")
	} else {
		t.Error("can not get resource of %s", resourceName)
	}
}

package main

import (
	"io/ioutil"
	"net/http"
)

func getPostBody(req *http.Request) ([]byte, error) {
	bytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	return bytes, err
}

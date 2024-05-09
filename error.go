package main

import (
	"fmt"
	"net/http"
)

func handleApiMethodError(writer http.ResponseWriter, req *http.Request, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	// todo: log stuff to logging framework
	http.Error(writer, message, http.StatusInternalServerError)
}

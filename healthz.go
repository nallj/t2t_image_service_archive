package main

import (
	"fmt"
	"net/http"
)

func handleHealthz(writer http.ResponseWriter) {
	// log.Print("I'm still here!")
	fmt.Fprintf(writer, "I'm still here!")
}

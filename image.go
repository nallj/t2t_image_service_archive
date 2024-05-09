package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "firebase.google.com/go/v4"
	// "firebase.google.com/go/v4/auth"
	// "firebase.google.com/go/v4/messaging"

	"github.com/nallj/t2t_image_service/config"
	"github.com/nallj/t2t_image_service/repository"
)

type ImageRequest struct {
	ImageId string `json:"imageId"`
}

func handleImage(config *config.Config, writer http.ResponseWriter, req *http.Request) {

	bytes, postError := getPostBody(req)
	if postError != nil {
		errMsg := fmt.Sprintf("Failed to read from request body: %s", postError.Error())
		log.Printf("Failed to read from request body: %s", errMsg)
		http.Error(writer, errMsg, 500)
		return
	}

	var request ImageRequest
	unmarshalError := json.Unmarshal(bytes, &request)
	if unmarshalError != nil {
		// TODO: Handle "Request body malformed: json: cannot unmarshal number into Go struct field ImageRequest.imageId of type string"
		errMsg := fmt.Sprintf("Request body malformed: %s", unmarshalError.Error())
		log.Printf(errMsg)
		http.Error(writer, errMsg, 500)
		return
	}

	// ImageId checks
	imageId := request.ImageId

	if imageId == "" {
		errMsg := "Empty image ID provided."
		log.Printf(errMsg)
		http.Error(writer, errMsg, 500)
		return
	}

	// TODO: Remove when working.
	log.Print("About to fetch from repository...")

	// TODO: Get JWT from request headers and verify they are a valid user.
	// req.Header().Get("bearer-token?")
	err := repository.GetImage(config, imageId, writer)
	if err != nil {
		log.Printf("Error occurred getting requested image: %v", err)
	}

	imageMsg := "Just testing things"
	// imageMsg := fmt.Sprintf("TODO: Got image: %s", image)
	log.Printf(imageMsg)
}

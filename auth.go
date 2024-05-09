package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/nallj/t2t_image_service/config"
	"github.com/nallj/t2t_image_service/repository"
)

type JwtRequest struct {
	UserId string `json:"userId"`
	Hash   string `json:"hash"`
}

const (
	ErrorUserFailAuth = "Provided user credentials are not valid."
)

func handleJwt(config *config.Config, writer http.ResponseWriter, req *http.Request) {

	bytes, err := getPostBody(req)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to read from request body: %s", err.Error())
		log.Printf("Failed to read from request body: %s", errMsg)
		http.Error(writer, errMsg, 500)
		return
	}

	var request JwtRequest
	err = json.Unmarshal(bytes, &request)
	if err != nil {
		errMsg := fmt.Sprintf("Request body malformed: %s", err.Error())
		log.Printf(errMsg)
		http.Error(writer, errMsg, 500)
		return
	}

	userId := request.UserId
	hash := request.Hash

	if userId == "" {
		errMsg := "Empty user ID provided."
		log.Printf(errMsg)
		http.Error(writer, errMsg, 500)
		return
	}

	if hash == "" {
		errMsg := "Empty hash provided."
		log.Printf(errMsg)
		http.Error(writer, errMsg, 500)
		return
	}

	authError := verifyUserAuth(userId, hash)
	if authError != nil {
		errMsg := authError.Error()
		log.Printf(
			"Mismatch between password for user ID %s and provided hash \"%s\"",
			userId,
			hash,
		)
		http.Error(writer, errMsg, 400)
		return
	}

	// TODO: Ok, you've got a valid user. Produce a JWT for them.
}

func verifyUserAuth(userId string, hash string) error {

	user := repository.GetUser(userId)
	assertError := assertHashPasswordMatch(hash, user.Password)
	if assertError != nil {
		return assertError
	}
	return nil // is this necessary?
}

func assertHashPasswordMatch(hash string, password string) error {

	hashedPassword := "" // TODO

	// Error if hashed password doesn't match provided hash.
	if hash != hashedPassword {
		return errors.New(ErrorUserFailAuth)
	}
	return nil // is this necessary?
}

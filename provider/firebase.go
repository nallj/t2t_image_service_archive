package provider

import (
	"context"
	"log"

	// "firebase.google.com/go/v4"
	// "firebase.google.com/go/v4/auth"
	// "firebase.google.com/go/v4/messaging"

	// ref: https://pkg.go.dev/cloud.google.com/go/firestore
	"cloud.google.com/go/firestore" // is it "github.com/googleapis/google-cloud-go"?
)

type User struct {
	Email    string
	Password string
}

func GetUserFromFirestore(userId string) (User, error) { // ??? { // TODO
	user := User{}
	return user, nil
}

// Returns a data record of with the target database record ID matching imageId.
// Returns if the the cooresponding file name if a match is made and any error that occurs.
func GetImageFromFirestore(imageId string) (string, error) { // TODO: Not really a string.

	// https://pkg.go.dev/cloud.google.com/go/firestore
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "t2t-app") //, "tow-2-tow-prototype")
	if err != nil {
		return "", err
	}

	collection := client.Collection("uploaded_image")
	imageDoc := collection.Doc(imageId)

	docSnapshot, err := imageDoc.Get(ctx)
	if err != nil {
		// rpc error: code = NotFound desc = "projects/t2t-app/databases/(default)/documents/uploaded_image/C6gVz0pTVwF8WWqxwInJ_" not found
		log.Printf("Error occurred while trying to fetch the record. %v", err)
		return "", err
	}

	exists := docSnapshot.Exists()
	if !exists {
		return "", nil
	}

	dataMap := docSnapshot.Data()
	fileName := dataMap["fileName"].(string)

	return fileName, nil
}

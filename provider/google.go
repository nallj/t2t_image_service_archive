package provider

import (
	"context"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/nallj/t2t_image_service/config"

	// for authentication? see https://pkg.go.dev/cloud.google.com/go#section-readme
	"google.golang.org/api/iterator"
)

const PROJECT_NAME string = "gothic-context-373302"
const BUCKET_NAME string = "t2t_business_images" // "gs://t2t_business_images"

// ref: https://cloud.google.com/go/docs/reference/cloud.google.com/go/storage/latest

// handling errors: https://cloud.google.com/go/docs/reference/cloud.google.com/go/storage/latest#hdr-Errors

func GetImageFromGoogle(cfg *config.Config, fileName string, writer http.ResponseWriter) error {

	ctx := context.Background()

	log.Print("Creating Google Cloud Storage client from library (without authentication).")

	// "Clients should be reused instead of created as needed. The methods of Client are safe for concurrent use by multiple goroutines."
	// client, err := storage.NewClient(ctx)
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
		log.Fatal(err)
	}
	defer client.Close()

	if client == nil {
		log.Print("client is nil")
	}

	log.Print("Client isn't nil. About to iterate over buckets.\n")

	it := client.Buckets(ctx, PROJECT_NAME)
	log.Print("Has iterator.\n")

	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Err happened. %v", err)
		}
		log.Print("About to print bucket name.\n")
		log.Printf("Bucket name. %v", bucketAttrs.Name)
	}

	log.Print("Listed all storage buckets.\n")

	log.Print("Getting bucket reference.")

	bucket := client.Bucket(BUCKET_NAME)

	log.Print("Getting bucket object reference.")
	log.Printf("THE FILESNAME: '%v'", fileName)

	// "Anonymous caller does not have storage.objects.get access to the Google Cloud Storage object."
	imageObj := bucket.Object(fileName)

	log.Print("Getting bucket object reader.")

	reader, err := imageObj.NewReader(ctx)

	// Error fetching bucket object: storage: object doesn't exist
	if err != nil {
		log.Printf("Error fetching bucket object: %v", err)
		return err
	}
	defer reader.Close()

	log.Print("Attempt to copy from reader.")

	// if _, err := io.Copy(os.Stdout, reader); err != nil {
	if _, err := io.Copy(writer, reader); err != nil {
		log.Printf("Error using bucket object reader: %v", err)
		return err
	}
	return nil
}

package repository

import (
	"net/http"

	"github.com/nallj/t2t_image_service/config"
	"github.com/nallj/t2t_image_service/provider"
)

func GetImage(config *config.Config, imageId string, writer http.ResponseWriter) error { // TODO: Not really a string.

	fileName, err := provider.GetImageFromFirestore(imageId)
	if err != nil {
		return err
	}

	err = provider.GetImageFromGoogle(config, fileName, writer)
	if err != nil {
		return err
	}
	return nil
}

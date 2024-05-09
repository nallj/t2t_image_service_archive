package repository

import (
	"github.com/nallj/t2t_image_service/provider" // Is this right?
)

func GetUser(userId string) provider.User { // TODO: Not really a string.
	user, err := provider.GetUserFromFirestore(userId)
	if err != nil {

	}
	return user
}

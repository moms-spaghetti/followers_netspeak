package storage

import (
	"github.com/moms-spaghetti/followers/internal/models"
)

func CreateMockUserStorage() []*models.User {
	return []*models.User{
		{
			ID: "fa89b65d-a900-43d9-8c95-69b4ecf567bc",
			Name: "John",
			FollowerIDs: []string{},
		},
	}
}
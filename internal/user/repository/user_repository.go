package repository

import models "grpc-user-serviceo/internal/model"

type UserRepository interface {
	GetUserByID(id int64) (*models.User, error)
	GetUsersByIDs(ids []int64) ([]models.User, error)
	SearchUsers(criteria map[string]interface{}) ([]models.User, error)
}

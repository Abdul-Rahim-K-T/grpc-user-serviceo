package memory

import (
	"errors"
	models "grpc-user-serviceo/internal/model"
	"grpc-user-serviceo/internal/user/repository"
)

type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []models.User{
			{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{ID: 2, FName: "Alice", City: "NY", Phone: 9876543210, Height: 5.5, Married: false},
		},
	}
}

func (r *UserRepository) GetUserByID(id int64) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found \n Pleas Enter valid")
}

func (r *UserRepository) GetUsersByIDs(ids []int64) ([]models.User, error) {
	var users []models.User
	for _, id := range ids {
		if user, err := r.GetUserByID(id); err == nil {
			users = append(users, *user)
		}
	}
	return users, nil
}

func (r *UserRepository) SearchUsers(criteria map[string]interface{}) ([]models.User, error) {
	// Implement search logic based on criteria
	var results []models.User
	for _, user := range r.users {
		matches := true
		for key, value := range criteria {
			switch key {
			case "city":
				if user.City != value {
					matches = false
				}
			case "phone":
				if user.Phone != value {
					matches = false
				}
			case "married":
				if user.Married != value {
					matches = false
				}
			}
		}
		if matches {
			results = append(results, user)
		}
	}
	return results, nil
}

// Ensure it satisfies the UserRepository interface
var _ repository.UserRepository = &UserRepository{}

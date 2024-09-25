package grpc

import (
	"context"
	"errors"
	"grpc-user-serviceo/grpc-user-serviceo/pkg/grpc/user"
	models "grpc-user-serviceo/internal/model"
	"testing"
)

type mockUserUsecase struct {
	users map[int64]*models.User
}

func (m *mockUserUsecase) FetchUserByID(id int64) (*models.User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func (m *mockUserUsecase) FetchUsersByIDs(ids []int64) ([]models.User, error) {
	var result []models.User
	for _, id := range ids {
		if user, ok := m.users[id]; ok {
			result = append(result, *user)
		}
	}
	return result, nil
}

func (m *mockUserUsecase) Search(criteria map[string]interface{}) ([]models.User, error) {
	var results []models.User
	for _, user := range m.users {
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
			results = append(results, *user)
		}
	}
	return results, nil
}

func TestUserHandler_GetUser(t *testing.T) {
	users := map[int64]*models.User{
		1: {ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	}
	usecase := &mockUserUsecase{users: users}
	handler := NewUserHandler(usecase)

	req := &user.UserIDRequest{Id: 1}
	got, err := handler.GetUser(context.Background(), req)

	if err != nil {
		t.Fatalf("GetUser() error = %v", err)
	}

	if got.User.Fname != "Steve" {
		t.Errorf("GetUser() = %v, want %v", got.User.Fname, "Steve")
	}
}

func TestUserHandler_GetUser_NotFound(t *testing.T) {
	usecase := &mockUserUsecase{users: map[int64]*models.User{}}
	handler := NewUserHandler(usecase)

	req := &user.UserIDRequest{Id: 2}
	_, err := handler.GetUser(context.Background(), req)

	if err == nil {
		t.Errorf("expected an error, got none")
	}
}

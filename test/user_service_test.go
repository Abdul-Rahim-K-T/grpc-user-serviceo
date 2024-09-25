package test

import (
	"context"
	"errors"
	pb "grpc-user-serviceo/grpc-user-serviceo/pkg/grpc/user" // Updated import
	models "grpc-user-serviceo/internal/model"
	userHandler "grpc-user-serviceo/internal/user/delivery/grpc"
	"grpc-user-serviceo/internal/user/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock Repository for testing
type mockUserRepository struct {
	users []models.User
}

func (m *mockUserRepository) GetUserByID(id int64) (*models.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *mockUserRepository) GetUsersByIDs(ids []int64) ([]models.User, error) {
	var result []models.User
	for _, id := range ids {
		for _, user := range m.users {
			if user.ID == id {
				result = append(result, user)
			}
		}
	}
	return result, nil
}

func (m *mockUserRepository) SearchUsers(criteria map[string]interface{}) ([]models.User, error) {
	var result []models.User
	for _, user := range m.users {
		if user.City == criteria["city"] {
			result = append(result, user)
		}
	}
	return result, nil
}

func setup() pb.UserServiceServer {
	users := []models.User{
		{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{ID: 2, FName: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
	}

	repo := &mockUserRepository{users: users}
	usecase := usecase.NewUserUsecase(repo)
	return userHandler.NewUserHandler(usecase) // Ensure this returns the correct type
}

func TestGetUserByID(t *testing.T) {
	handler := setup()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test panicked: %v", r)
		}
	}()

	req := &pb.UserIDRequest{Id: 1}
	resp, err := handler.GetUser(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(1), resp.User.Id)
	assert.Equal(t, "Steve", resp.User.Fname)
}

func TestGetUserByID_NotFound(t *testing.T) {
	handler := setup()

	req := &pb.UserIDRequest{Id: 99}
	resp, err := handler.GetUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "user not found", err.Error())
}

func TestSearchUsers(t *testing.T) {
	handler := setup()

	req := &pb.SearchRequest{
		City: "LA",
	}

	resp, err := handler.SearchUsers(context.Background(), req)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test panicked: %v", r)
		}
	}()

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp.Users))
	assert.Equal(t, "Steve", resp.Users[0].Fname)
}

func TestSearchUsers_NoResults(t *testing.T) {
	handler := setup()

	req := &pb.SearchRequest{
		City: "Miami",
	}

	resp, err := handler.SearchUsers(context.Background(), req)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test panicked: %v", r)
		}
	}()

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 0, len(resp.Users))
}

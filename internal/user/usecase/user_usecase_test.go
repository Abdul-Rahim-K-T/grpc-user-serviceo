package usecase

import (
	"errors"
	models "grpc-user-serviceo/internal/model"
	"testing"
)

type mockRepo struct {
	users map[int64]*models.User
}

func (m *mockRepo) GetUserByID(id int64) (*models.User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found ")
}

func (m *mockRepo) GetUsersByIDs(ids []int64) ([]models.User, error) {
	var result []models.User
	for _, id := range ids {
		if user, ok := m.users[id]; ok {
			result = append(result, *user)
		}
	}
	return result, nil
}

func (m *mockRepo) SearchUsers(criteria map[string]interface{}) ([]models.User, error) {
	var results []models.User
	for _, user := range m.users {
		matches := true
		for key, value := range criteria {
			switch key {
			case "City":
				if user.City != value {
					matches = false
				}
			case "Married":
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

func TestUserUsecase_FetchUserByID(t *testing.T) {
	users := map[int64]*models.User{
		1: {ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	}
	repo := &mockRepo{users: users}
	usecase := NewUserUsecase(repo)

	got, err := usecase.FetchUserByID(1)
	if err != nil {
		t.Fatalf("FetchUserByID() error = %v", err)
	}
	if got.ID != 1 {
		t.Errorf("FetchUserByID() = %v, want %v", got.ID, 1)
	}
}

func TestUserUsecase_Search(t *testing.T) {
	users := map[int64]*models.User{
		1: {ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		2: {ID: 2, FName: "Alice", City: "NY", Phone: 987654321, Height: 5.5, Married: false},
	}
	repo := &mockRepo{users: users}
	usecase := NewUserUsecase(repo)

	tests := []struct {
		criteria map[string]interface{}
		wantLen  int
	}{
		{map[string]interface{}{"City": "LA"}, 1},
		{map[string]interface{}{"Married": true}, 1},
	}

	for _, tt := range tests {
		t.Run("Search", func(t *testing.T) {
			got, err := usecase.Search(tt.criteria)
			if err != nil {
				t.Fatalf("Search() error = %v", err)
			}
			if len(got) != tt.wantLen {
				t.Errorf("Search() = %v want length %v", len(got), tt.wantLen)
			}
		})
	}
}

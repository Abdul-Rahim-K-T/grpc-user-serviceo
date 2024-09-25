package memory

import (
	models "grpc-user-serviceo/internal/model"
	"testing"
)

func TestUserRepository_GetUserByID(t *testing.T) {
	repo := NewUserRepository()

	tests := []struct {
		id      int64
		want    *models.User
		wantErr bool
	}{
		{1, &models.User{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}, false},
		{3, nil, true},
	}

	for _, tt := range tests {
		t.Run("GetUserByID", func(t *testing.T) {
			got, err := repo.GetUserByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.ID != tt.want.ID {
				t.Errorf("GetUserByID() %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetUsersByIDs(t *testing.T) {
	repo := NewUserRepository()

	tests := []struct {
		ids     []int64
		wantLen int
	}{
		{[]int64{1, 2}, 2},
		{[]int64{1, 3}, 1},
	}

	for _, tt := range tests {
		t.Run("GetUserByIDs", func(t *testing.T) {
			got, err := repo.GetUsersByIDs(tt.ids)
			if err != nil {
				t.Errorf("GetUsersByIDs() error = %v", err)
			}
			if len(got) != tt.wantLen {
				t.Errorf("GetUsersByIDs() = %v, want length %v", len(got), tt.wantLen)
			}
		})
	}
}

func TestUserRepository_SearchUsers(t *testing.T) {
	repo := NewUserRepository()

	tests := []struct {
		name     string
		criteria map[string]interface{}
		want     []models.User
	}{
		{
			name:     "Search by city LA",
			criteria: map[string]interface{}{"city": "LA"},
			want: []models.User{
				{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			},
		},
		{
			name:     "Search by phone",
			criteria: map[string]interface{}{"phone": int64(9876543210)},
			want: []models.User{
				{ID: 2, FName: "Alice", City: "NY", Phone: 9876543210, Height: 5.5, Married: false},
			},
		},
		{
			name:     "Search by married status",
			criteria: map[string]interface{}{"married": true},
			want: []models.User{
				{ID: 1, FName: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			},
		},
		{
			name:     "Search by city and married status",
			criteria: map[string]interface{}{"city": "NY", "married": false},
			want: []models.User{
				{ID: 2, FName: "Alice", City: "NY", Phone: 9876543210, Height: 5.5, Married: false},
			},
		},
		{
			name:     "No matches",
			criteria: map[string]interface{}{"city": "SF"},
			want:     []models.User{},
		},
	}

	for _, tt := range tests {
		t.Run("SearchUsers", func(t *testing.T) {
			got, err := repo.SearchUsers(tt.criteria)
			if err != nil {
				t.Errorf("SearchUsers() error = %v", err)
				return
			}
			if !compareUsers(got, tt.want) {
				t.Errorf("SearchUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareUsers(a, b []models.User) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].ID != b[i].ID || a[i].FName != b[i].FName || a[i].City != b[i].City || a[i].Phone != b[i].Phone || a[i].Height != b[i].Height || a[i].Married != b[i].Married {
			return false
		}
	}
	return true
}

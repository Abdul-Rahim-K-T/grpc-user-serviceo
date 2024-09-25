package usecase

import (
	models "grpc-user-serviceo/internal/model"
	"grpc-user-serviceo/internal/user/repository"
)

type UserUsecase interface {
	FetchUserByID(id int64) (*models.User, error)
	FetchUsersByIDs(ids []int64) ([]models.User, error)
	Search(criteria map[string]interface{}) ([]models.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) FetchUserByID(id int64) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *userUsecase) FetchUsersByIDs(ids []int64) ([]models.User, error) {
	return u.repo.GetUsersByIDs(ids)
}

func (u *userUsecase) Search(criteria map[string]interface{}) ([]models.User, error) {
	return u.repo.SearchUsers(criteria)
}

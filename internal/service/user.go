package service

import (
	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/internal/repo"
)

type UserService struct {
	repo repo.User
}

func NewUserService(repo repo.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(user entity.User) (int, error) {
	return u.repo.CreateUser(user)
}

func (u *UserService) GetUser(id int) (entity.User, error) {
	return u.repo.GetUser(id)
}

func (u *UserService) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}

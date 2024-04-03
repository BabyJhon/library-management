package service

import (
	"context"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/internal/repo"
)

type UserService struct {
	repo repo.User
}

func NewUserService(repo repo.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(ctx context.Context, user entity.User) (int, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserService) GetUser(ctx context.Context, id int) (entity.User, error) {
	return u.repo.GetUser(ctx, id)
}

func (u *UserService) DeleteUser(ctx context.Context, id int) error {
	return u.repo.DeleteUser(ctx, id)
}

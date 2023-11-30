package storage

import (
	"context"
	"user/models"
)

type StorageI interface {
	User() UsersI
}

type UsersI interface {
	CreateUser(context.Context, *models.CreateUser) (string, error)
	GetUser(context.Context, *models.IdRequest) (*models.User, error)
	GetAllUser(context.Context, *models.GetAllUserRequest) (*models.GetAllUser, error)
	UpdateUser(context.Context, *models.UpdateUser) (string, error)
	DeleteUser(context.Context, *models.IdRequest) (string, error)
}

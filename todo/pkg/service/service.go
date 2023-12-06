package service

import (
	models "todo/models"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user *models.CreateUser) (int, error)
	GetUserById(req *models.IdRequest) (rep models.CreateUser, err error)
	GetAllUsers(req *models.GetAllUserRequest) (rep models.GetAllUser, err error)
	UpdateUser(req *models.User) (string, error)
	DeleteUser(req *models.IdRequest) (string, error)
	CreateUsers(user []models.CreateUser) ([]int, error)
	UpdateUsers(req []models.User) ([]string, error)

	// GenerateToken(username, password string) (string, error)
	// ParseToken(token string) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Authorization
	// TodoList
	// TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{

		Authorization: NewAuthService(repos.Authorization),
	}
}

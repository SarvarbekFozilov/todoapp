package repository

import (
	models "todo/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user *models.CreateUser) (int, error)
	GetUserById(req *models.IdRequest) (rep models.CreateUser, err error)
	GetAllUsers(req *models.GetAllUserRequest) (rep models.GetAllUser, err error)
	UpdateUser(req *models.User) (string, error)
	DeleteUser(req *models.IdRequest) (string, error)
	CreateUsers(user []models.CreateUser) ([]int, error)
	UpdateUsers(req []models.User) ([]string, error)
}
type TodoList interface {
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgress(db),
	}
}

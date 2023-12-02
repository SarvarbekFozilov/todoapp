package repository

import (
	models "todo/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(req *models.CreateUser) (int, error)
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

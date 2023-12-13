package repository

import (
	models "todo/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user *models.CreateUserReq) (int, error)
	GetUser(username , password string) ( *models.UserResponse,  error)

}
type User interface {
	CreateUser(user *models.CreateUserReq) (int, error)
	GetUserById(req *models.IdRequest) (rep* models.UserResponse, err error)
	GetAllUsers(req *models.GetAllUserReq) (rep models.GetAllUser, err error)
	UpdateUser(req *models.UpdateUser) (int, error)
	DeleteUser(req *models.IdRequest) (int, error)
	CreateUsers(user []models.CreateUserReq) ([]int, error)
	UpdateUsers(users []models.UpdateUser) (string, error)

}


type Repository struct {
	Authorization
	User

}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgress(db),
		User:NewUserPostgress(db),
	}
}

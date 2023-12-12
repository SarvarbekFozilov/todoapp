package service

import (
	models "todo/models"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user *models.CreateUserReq) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)


}
type User interface {
	CreateUser(user *models.CreateUserReq) (int, error)
	GetUserById(req *models.IdRequest) (rep *models.UserResponse, err error)
	GetAllUsers(req *models.GetAllUserReq) (rep models.GetAllUser, err error)
	UpdateUser(req *models.UpdateUser) (int, error)
	DeleteUser(req *models.IdRequest) (int, error)
	CreateUsers(user []models.CreateUserReq) ([]int, error)
	UpdateUsers(req []models.UpdateUser) ([]int, error)

}

type Service struct {
	Authorization
	User

}


func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:NewUserService(repos.User),
	}
}

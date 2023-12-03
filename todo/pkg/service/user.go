package service

import (
	"crypto/sha1"
	"fmt"
	models "todo/models"
	"todo/pkg/repository"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CreateUser(user *models.CreateUser) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserById(req *models.IdRequest) (rep models.CreateUser, err error) {
	return s.repo.GetUserById(req)
}
func (s *AuthService) GetAllUsers(req *models.GetAllUserRequest) (rep models.GetAllUser, err error) {
	return s.repo.GetAllUsers(req)
}

func (s *AuthService) UpdateUser(req *models.User) (string, error) {

	return s.repo.UpdateUser(req)
}
func (s *AuthService) DeleteUser(req *models.IdRequest) (string, error) {
	return s.repo.DeleteUser(req)
}
func (s *AuthService) CreateUsers(users []models.CreateUser) ([]int, error) {
	for i := range users {
		users[i].Password = generatePasswordHash(users[i].Password)
	}

	return s.repo.CreateUsers(users)
}

func (s *AuthService) UpdateUsers(req []models.User) (string, error) {

	return s.repo.UpdateUsers(req)
}

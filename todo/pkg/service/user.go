package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	models "todo/models"
	"todo/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.CreateUserReq) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserById(req *models.IdRequest) (rep *models.UserResponse, err error) {
	return s.repo.GetUserById(req)
}
func (s *UserService) GetAllUsers(req *models.GetAllUserReq) (rep models.GetAllUser, err error) {
	return s.repo.GetAllUsers(req)
}

func (s *UserService) UpdateUser(req *models.UpdateUser) (int, error) {

	return s.repo.UpdateUser(req)
}
func (s *UserService) DeleteUser(req *models.IdRequest) (int, error) {
	return s.repo.DeleteUser(req)
}

func (s *UserService) CreateUsers(users []models.CreateUserReq) ([]int, error) {
	for i := range users {
		users[i].Password = generatePasswordHash(users[i].Password)
	}

	return s.repo.CreateUsers(users)
}
func (s *UserService) UpdateUsers(users []models.UpdateUser) (string, error) {
	for i := range users {
		users[i].Password = generatePasswordHash(users[i].Password)
	}

	res, err := s.repo.UpdateUsers(users)
	if err != nil {
		fmt.Printf("service: %v", err)
		return "", err
	}

	return res, nil
}

func convertStringIDsToInts(ids string) ([]int, error) {
	idStrings := strings.Split(ids, ",")
	var result []int

	for _, idStr := range idStrings {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, errors.New("invalid user ID")
		}
		result = append(result, id)
	}

	return result, nil
}

func convertToStrings(fields string) ([]string, error) {
	fieldStrings := strings.Split(fields, ",")
	var result []string

	result = append(result, fieldStrings...)

	return result, nil
}

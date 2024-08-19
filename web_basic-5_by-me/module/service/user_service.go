package service

import (
	"user_server/model"
	"user_server/module/repository"
)

type UserService interface {
	Create(user *model.User) error
	GetAllUsers() ([]model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// --------------Implement 物件-------------------------

type userService struct {
	repo repository.UserRepository // 如此可以讓 userService method 使用 repository 中的 method了 (使用此 interface)
}

func (s *userService) Create(user *model.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserById(id int) (model.User, error) {
	return s.repo.GetById(id)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(user *model.User) error {
	_, err := s.repo.GetById(int(user.ID))
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}
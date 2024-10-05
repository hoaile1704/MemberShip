package services

import repository "UserService/internal/repository"

type UserService interface {
	GetUser(id uint) (*repository.User, error)
	CreateUser(name string, age int) (*repository.User, error)
}

type UserServiceImpl struct {
	UserRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepo: repo}
}

func (s *UserServiceImpl) GetUser(id uint) (*repository.User, error) {
	return s.UserRepo.FindByID(id)
}

func (s *UserServiceImpl) CreateUser(name string, age int) (*repository.User, error) {
	user := &repository.User{Name: name, Age: age}
	err := s.UserRepo.Save(user)
	return user, err
}

package repository

import "gorm.io/gorm"

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
	Age  int    `gorm:"not null"`
}

type UserRepository interface {
	FindByID(id uint) (*User, error)
	Save(user *User) error
}

type UserServiceImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserServiceImpl{DB: db}
}

func (r *UserServiceImpl) FindByID(id uint) (*User, error) {
	var user User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserServiceImpl) Save(user *User) error {
	return r.DB.Create(user).Error
}

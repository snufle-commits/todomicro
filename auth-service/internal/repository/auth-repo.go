package repository

import (
	"authservice/internal/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	SignUp(user *model.User) error // исправлено: принимаем указатель
	SignIn(email string) (*model.User, error)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepository {
	return &authRepo{db: db} // исправлено: возвращаем указатель и правильный тип
}

func (r *authRepo) SignUp(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *authRepo) SignIn(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

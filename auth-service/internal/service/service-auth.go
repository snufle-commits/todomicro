package service

import (
	"authservice/internal/model"
	"authservice/internal/repository"
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var SECRET string = "gfdsgsfgfsg"
var TKN string

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(email string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{Email: email, Password: string(hash)}

	return s.repo.SignUp(&user)
}

func (s *AuthService) SignIn(email, password string) (string, error) {
	user, err := s.repo.SignIn(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	TokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	TKN = TokenString
	return TKN, nil
}

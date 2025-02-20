package auth

import (
	"fmt"

	"github.com/holycann/smart-parking-backend/config"
	"github.com/holycann/smart-parking-backend/internal/middleware"
	"github.com/holycann/smart-parking-backend/internal/users"
)

type UserService struct {
	repository *users.UserRepository
}

func NewService(repository *users.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) UserLogin(payload *LoginUserPayload) (string, error) {
	u, err := s.repository.GetUserByEmail(payload.Email)
	if err != nil {
		return "", fmt.Errorf("error get user by email: %v\n", err)
	}

	if !middleware.ComparePassword(u.Password, []byte(payload.Password)) {
		return "", fmt.Errorf("error compare password: %v\n", err)
	}

	token, err := middleware.CreateJWT([]byte(config.Env.JWTSecret), u.ID)
	if err != nil {
		return "", fmt.Errorf("error create jwt: %v\n", err)
	}

	return token, nil
}

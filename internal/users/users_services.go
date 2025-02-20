package users

import (
	"fmt"

	"github.com/holycann/smart-parking-backend/internal/middleware"
)

type UserService struct {
	repository UserRepositoryInterface
}

func NewService(repository UserRepositoryInterface) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetAllUserData() ([]*User, error) {
	return s.repository.GetAllUser()
}

func (s *UserService) GetUserByID(id int) (*User, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("error getting user by id: %v\n", err)
	}

	return user, nil
}

func (s *UserService) CreateUser(payload *CreateUserPayload) (string, error) {
	_, err := s.repository.GetUserByEmail(payload.Email)
	if err == nil {
		return "", fmt.Errorf("Email %s already exists", payload.Email)
	}

	hashedPassword, err := middleware.HashPassword(payload.Password)
	if err != nil {
		return "", fmt.Errorf("Failed To Hash Password: %v", err)
	}

	err = s.repository.CreateUser(&CreateUserPayload{
		Fullname:    payload.Fullname,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
		Password:    hashedPassword,
		ImageURL:    payload.ImageURL,
	})
	if err != nil {
		return "", fmt.Errorf("error create user: %v\n", err)
	}

	return fmt.Sprintf("Create user %s successfully", payload.Fullname), nil
}

func (s *UserService) UpdateUser(payload *UpdateUserPayload) (string, error) {
	u, err := s.repository.GetUserByID(payload.ID)
	if err != nil {
		return "", fmt.Errorf("error get user by id: %v\n", err)
	}

	if u == nil {
		return "", fmt.Errorf("User with ID %d does not exist", payload.ID)
	}

	if payload.Fullname == "" && payload.PhoneNumber == "" {
		return "", fmt.Errorf("User Fullname And Phone Number Cannot Be Empty!")
	}

	u, err = s.repository.GetUserByEmail(payload.Email)
	if err != nil {
		return "", fmt.Errorf("error get user by username: %v\n", err)
	}

	err = s.repository.UpdateUser(&UpdateUserPayload{
		ID:          payload.ID,
		Fullname:    payload.Fullname,
		PhoneNumber: payload.PhoneNumber,
		ImageURL:    payload.ImageURL,
	})
	if err != nil {
		return "", fmt.Errorf("error update user: %v\n", err)
	}

	return fmt.Sprintf("Update user %s successfully", u.Fullname), nil
}

func (s *UserService) DeleteUser(id int) (string, error) {
	err := s.repository.DeleteUser(id)
	if err != nil {
		return "", fmt.Errorf("error delete user: %v\n", err)
	}

	return fmt.Sprintf("Delete user successfully"), nil
}

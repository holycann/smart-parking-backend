package users

import "database/sql"

type UserRepositoryInterface interface {
	GetAllUser() ([]*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *CreateUserPayload) error
	UpdateUser(user *UpdateUserPayload) error
	DeleteUser(id int) error
}

type UserServiceInterface interface {
	GetAllUserData() ([]*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(payload *CreateUserPayload) (string, error)
	UpdateUser(payload *UpdateUserPayload) (string, error)
	DeleteUser(id int) (string, error)
}

type User struct {
	ID          int          `json:"id"`
	Fullname    string       `json:"fullname"`
	Email       string       `json:"email"`
	PhoneNumber string       `json:"phone_number"`
	Password    string       `json:"password"`
	ImageURL    string       `json:"image_url"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type CreateUserPayload struct {
	Fullname    string `json:"fullname" validate:"required,min=3,max=50"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=20"`
	Password    string `json:"password" validate:"required,min=8,max=32"`
	ImageURL    string `json:"image_url"`
}

type UpdateUserPayload struct {
	ID          int    `json:"id" validate:"required"`
	Fullname    string `json:"fullname" validate:"required,min=3,max=30"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=20"`
	Password    string `json:"password" validate:"required,min=8,max=32"`
	ImageURL    string `json:"image_url"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

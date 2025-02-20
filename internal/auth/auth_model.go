package auth

type AuthServiceInterface interface {
	UserLogin(user *LoginUserPayload) (string, error)
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

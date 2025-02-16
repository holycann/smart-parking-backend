package payment_methods

import (
	"database/sql"
)

type PaymentMethodStore interface {
	GetAllPaymentMethod() ([]*PaymentMethod, error)
	GetPaymentMethodByMethodName(MethodName string) (*PaymentMethod, error)
	GetPaymentMethodByID(id int) (*PaymentMethod, error)
	CreatePaymentMethod(payment_method *CreatePaymentMethodPayload) error
	UpdatePaymentMethod(payment_method *UpdatePaymentMethodPayload) error
	DeletePaymentMethod(id int) error
}

type PaymentMethod struct {
	ID         int          `json:"id"`
	MethodName string       `json:"method_name"`
	Details    string       `json:"details"`
	Status     string       `json:"status"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type CreatePaymentMethodPayload struct {
	MethodName string `json:"method_name" validate:"required"`
	Details    string `json:"details" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

type UpdatePaymentMethodPayload struct {
	ID         int    `json:"id" validate:"required"`
	MethodName string `json:"method_name" validate:"required"`
	Details    string `json:"details" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

package transactions

import "database/sql"

type TransactionRepositoryInterface interface {
	GetAllTransaction() ([]*Transaction, error)
	GetTransactionByReservationID(ReservationID int) (*Transaction, error)
	GetTransactionByID(id int) (*Transaction, error)
	CreateTransaction(transaction *CreateTransactionPayload) error
	UpdateTransaction(transaction *UpdateTransactionPayload) error
	DeleteTransaction(id int) error
}

type TransactionServiceInterface interface {
	GetAllTransaction() ([]*Transaction, error)
	GetTransactionByID(id int) (*Transaction, error)
	CreateTransaction(transaction *CreateTransactionPayload) (string, error)
	UpdateTransaction(transaction *UpdateTransactionPayload) (string, error)
	DeleteTransaction(id int) (string, error)
}

type Transaction struct {
	ID              int          `json:"id"`
	ReservationID   int          `json:"reservation_id"`
	Amount          int          `json:"amount"`
	PaymentMethodID int          `json:"payment_method_id"`
	Status          string       `json:"status"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

type CreateTransactionPayload struct {
	ReservationID   int    `json:"reservation_id" validate:"required"`
	Amount          int    `json:"amount" validate:"required"`
	PaymentMethodID int    `json:"payment_method_id" validate:"required"`
	Status          string `json:"status" validate:"required"`
}

type UpdateTransactionPayload struct {
	ID              int    `json:"id" validate:"required"`
	ReservationID   int    `json:"reservation_id" validate:"required"`
	Amount          int    `json:"amount" validate:"required"`
	PaymentMethodID int    `json:"payment_method_id" validate:"required"`
	Status          string `json:"status" validate:"required"`
}

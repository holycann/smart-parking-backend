package transactions

import (
	"fmt"
)

type TransactionService struct {
	repository TransactionRepositoryInterface
}

func NewService(repository TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{repository: repository}
}

func (s *TransactionService) GetAllTransaction() ([]*Transaction, error) {
	return s.repository.GetAllTransaction()
}

func (s *TransactionService) GetTransactionByID(id int) (*Transaction, error) {
	transaction, err := s.repository.GetTransactionByID(id)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("error getting transaction by id: %v\n", err)
	}

	return transaction, nil
}

func (s *TransactionService) CreateTransaction(payload *CreateTransactionPayload) (string, error) {
	_, err := s.repository.GetTransactionByReservationID(payload.ReservationID)
	if err == nil {
		return "", fmt.Errorf("Transaction Number %s already exists", payload.ReservationID)
	}

	err = s.repository.CreateTransaction(&CreateTransactionPayload{
		ReservationID:   payload.ReservationID,
		Amount:          payload.Amount,
		PaymentMethodID: payload.PaymentMethodID,
		Status:          payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error create transaction: %v\n", err)
	}

	return fmt.Sprintf("Create transaction %s successfully", payload.ReservationID), nil
}

func (s *TransactionService) UpdateTransaction(payload *UpdateTransactionPayload) (string, error) {
	t, err := s.repository.GetTransactionByID(payload.ID)
	if err != nil {
		return "", fmt.Errorf("error get transaction by id: %v\n", err)
	}

	if t == nil {
		return "", fmt.Errorf("Transaction wits ID %d does not exist", payload.ID)
	}

	if payload.ReservationID == 0 && payload.PaymentMethodID == 0 {
		return "", fmt.Errorf("Transaction Reservation And Payment Method Cannot Be Empty!")
	}

	err = s.repository.UpdateTransaction(&UpdateTransactionPayload{
		ID:              payload.ID,
		ReservationID:   payload.ReservationID,
		Amount:          payload.Amount,
		PaymentMethodID: payload.PaymentMethodID,
		Status:          payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error update transaction: %v\n", err)
	}

	return fmt.Sprintf("Update transaction %s successfully", t.PaymentMethodID), nil
}

func (s *TransactionService) DeleteTransaction(id int) (string, error) {
	err := s.repository.DeleteTransaction(id)
	if err != nil {
		return "", fmt.Errorf("error delete transaction: %v\n", err)
	}

	return fmt.Sprintf("Delete transaction successfully"), nil
}

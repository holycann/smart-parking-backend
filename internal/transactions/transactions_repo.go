package transactions

import (
	"database/sql"
	"fmt"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func scanRowIntoTransaction(row *sql.Rows) (*Transaction, error) {
	transaction := new(Transaction)

	err := row.Scan(
		&transaction.ID,
		&transaction.ReservationID,
		&transaction.Amount,
		&transaction.PaymentMethodID,
		&transaction.Status,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *TransactionRepository) GetAllTransaction() (transaction []*Transaction, err error) {
	rows, err := s.db.Query("SELECT * FROM transactions")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		sp, err := scanRowIntoTransaction(rows)
		if err != nil {
			return nil, err
		}
		transaction = append(transaction, sp)
	}

	return transaction, nil
}

func (s *TransactionRepository) GetTransactionByReservationID(ReservationID int) (*Transaction, error) {
	rows, err := s.db.Query("SELECT * FROM transactions WHERE reservation_id = $1", ReservationID)
	if err != nil {
		return nil, err
	}

	t := new(Transaction)
	for rows.Next() {
		t, err = scanRowIntoTransaction(rows)
		if err != nil {
			return nil, err
		}
	}

	if t.ReservationID == 0 {
		return nil, fmt.Errorf("Transaction Reservation ID number not found")
	}

	return t, nil
}

func (s *TransactionRepository) GetTransactionByID(id int) (*Transaction, error) {
	rows, err := s.db.Query("SELECT * FROM transactions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	t := new(Transaction)
	for rows.Next() {
		t, err = scanRowIntoTransaction(rows)
		if err != nil {
			return nil, err
		}
	}

	if t.ID == 0 {
		return nil, fmt.Errorf("Transaction not found")
	}

	return t, nil
}

func (s *TransactionRepository) CreateTransaction(transaction *CreateTransactionPayload) error {
	_, err := s.db.Exec("INSERT INTO transactions (reservation_id, amount, payment_method_id, status) VALUES ($1, $2, $3 $4)", transaction.ReservationID, transaction.Amount, transaction.PaymentMethodID, transaction.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *TransactionRepository) UpdateTransaction(transaction *UpdateTransactionPayload) error {
	_, err := s.db.Exec("UPDATE transactions SET reservation_id = $1, amount = $2, payment_method_id = $3, status = $4 WHERE id = $5", transaction.ReservationID, transaction.Amount, transaction.PaymentMethodID, transaction.Status, transaction.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TransactionRepository) DeleteTransaction(id int) error {
	_, err := s.db.Exec("DELETE FROM transactions WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

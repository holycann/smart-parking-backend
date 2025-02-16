package payment_methods

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func scanRowIntoPaymentMethod(row *sql.Rows) (*PaymentMethod, error) {
	payment_method := new(PaymentMethod)

	err := row.Scan(
		&payment_method.ID,
		&payment_method.MethodName,
		&payment_method.Details,
		&payment_method.Status,
		&payment_method.CreatedAt,
		&payment_method.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return payment_method, nil
}

func (s *Store) GetAllPaymentMethod() (payment_method []*PaymentMethod, err error) {
	rows, err := s.db.Query("SELECT * FROM payment_methods")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		sp, err := scanRowIntoPaymentMethod(rows)
		if err != nil {
			return nil, err
		}
		payment_method = append(payment_method, sp)
	}

	return payment_method, nil
}

func (s *Store) GetPaymentMethodByMethodName(MethodName string) (*PaymentMethod, error) {
	rows, err := s.db.Query("SELECT * FROM payment_methods WHERE method_name = $1", MethodName)
	if err != nil {
		return nil, err
	}

	pm := new(PaymentMethod)
	for rows.Next() {
		pm, err = scanRowIntoPaymentMethod(rows)
		if err != nil {
			return nil, err
		}
	}

	if pm.MethodName == "" {
		return nil, fmt.Errorf("Payment Method Name not found")
	}

	return pm, nil
}

func (s *Store) GetPaymentMethodByID(id int) (*PaymentMethod, error) {
	rows, err := s.db.Query("SELECT * FROM payment_methods WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	pm := new(PaymentMethod)
	for rows.Next() {
		pm, err = scanRowIntoPaymentMethod(rows)
		if err != nil {
			return nil, err
		}
	}

	if pm.ID == 0 {
		return nil, fmt.Errorf("PaymentMethod not found")
	}

	return pm, nil
}

func (s *Store) CreatePaymentMethod(payment_method *CreatePaymentMethodPayload) error {
	_, err := s.db.Exec("INSERT INTO payment_methods (method_name, details, status) VALUES ($1, $2, $3", payment_method.MethodName, payment_method.Details, payment_method.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdatePaymentMethod(payment_method *UpdatePaymentMethodPayload) error {
	_, err := s.db.Exec("UPDATE payment_methods SET method_name = $1, details = $2, status = $3 WHERE id = $4", payment_method.MethodName, payment_method.Details, payment_method.Status, payment_method.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeletePaymentMethod(id int) error {
	_, err := s.db.Exec("DELETE FROM payment_methods WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

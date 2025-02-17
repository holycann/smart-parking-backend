package payment_methods

import (
	"fmt"
)

type PaymentMethodService struct {
	repository PaymentMethodRepositoryInterface
}

func NewService(repository PaymentMethodRepositoryInterface) *PaymentMethodService {
	return &PaymentMethodService{repository: repository}
}

func (s *PaymentMethodService) GetAllPaymentMethod() ([]*PaymentMethod, error) {
	return s.repository.GetAllPaymentMethod()
}

func (s *PaymentMethodService) GetPaymentMethodByID(id int) (*PaymentMethod, error) {
	payment, err := s.repository.GetPaymentMethodByID(id)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("error getting payment_method by id: %v\n", err)
	}

	return payment, nil
}

func (s *PaymentMethodService) CreatePaymentMethod(payload *CreatePaymentMethodPayload) (string, error) {
	_, err := s.repository.GetPaymentMethodByMethodName(payload.MethodName)
	if err == nil {
		return "", fmt.Errorf("Payment Method With Method Name %s already exists", payload.MethodName)
	}

	err = s.repository.CreatePaymentMethod(&CreatePaymentMethodPayload{
		MethodName: payload.MethodName,
		Details:    payload.Details,
		Status:     payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error create payment_method: %v\n", err)
	}

	return fmt.Sprintf("Create payment_method %s successfully", payload.MethodName), nil
}

func (s *PaymentMethodService) UpdatePaymentMethod(payload *UpdatePaymentMethodPayload) (string, error) {
	pm, err := s.repository.GetPaymentMethodByID(payload.ID)
	if err != nil {
		return "", fmt.Errorf("error get payment_method by id: %v\n", err)
	}
	if pm == nil {
		return "", fmt.Errorf("Payment_method with ID %d does not exist", payload.ID)
	}

	if payload.MethodName == "" {
		return "", fmt.Errorf("Payment Method Name Cannot Be Empty!!!")
	}

	err = s.repository.UpdatePaymentMethod(&UpdatePaymentMethodPayload{
		ID:         payload.ID,
		MethodName: payload.MethodName,
		Details:    payload.Details,
		Status:     payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error update payment: %v\n", err)
	}

	return fmt.Sprintf("Update payment method %v successfully", pm.MethodName), nil
}

func (s *PaymentMethodService) DeletePaymentMethod(id int) (string, error) {
	err := s.repository.DeletePaymentMethod(id)
	if err != nil {
		return "", fmt.Errorf("error delete payment_method: %v\n", err)
	}

	return fmt.Sprintf("Delete payment_method successfully"), nil
}

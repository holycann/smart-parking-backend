package notifications

import (
	"fmt"
)

type NotificationService struct {
	repository NotificationRepositoryInterface
}

func NewService(repository NotificationRepositoryInterface) *NotificationService {
	return &NotificationService{repository: repository}
}

func (s *NotificationService) GetAllNotification() ([]*Notification, error) {
	return s.repository.GetAllNotification()
}

func (s *NotificationService) GetNotificationByID(id int) (*Notification, error) {
	notification, err := s.repository.GetNotificationByID(id)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("error getting notification by id: %v\n", err)
	}

	return notification, nil
}

func (s *NotificationService) CreateNotification(payload *CreateNotificationPayload) (string, error) {
	_, err := s.repository.GetNotificationByMessage(payload.Message)
	if err == nil {
		return "", fmt.Errorf("Notification Message %s already exists", payload.Message)
	}

	err = s.repository.CreateNotification(&CreateNotificationPayload{
		UserID:  payload.UserID,
		Message: payload.Message,
		Status:  payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error create notification: %v\n", err)
	}

	return fmt.Sprintf("Create notification %s successfully", payload.Message), nil
}

func (s *NotificationService) UpdateNotification(payload *UpdateNotificationPayload) (string, error) {
	n, err := s.repository.GetNotificationByID(payload.ID)
	if err != nil {
		return "", fmt.Errorf("error get notification by id: %v\n", err)
	}

	if n == nil {
		return "", fmt.Errorf("Notification with ID %d does not exist", payload.ID)
	}

	if payload.UserID == 0 && payload.Message == "" {
		return "", fmt.Errorf("Notification User ID And Message Cannot Be Empty!")
	}

	err = s.repository.UpdateNotification(&UpdateNotificationPayload{
		ID:      payload.ID,
		UserID:  payload.UserID,
		Message: payload.Message,
		Status:  payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error update notification: %v\n", err)
	}

	return fmt.Sprintf("Update notification %s successfully", n.Message), nil
}

func (s *NotificationService) DeleteNotification(id int) (string, error) {
	err := s.repository.DeleteNotification(id)
	if err != nil {
		return "", fmt.Errorf("error delete notification: %v\n", err)
	}

	return fmt.Sprintf("Delete notification successfully"), err
}

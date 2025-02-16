package notifications

import "database/sql"

type NotificationStore interface {
	GetAllNotification() ([]*Notification, error)
	GetNotificationByMessage(message string) (*Notification, error)
	GetNotificationByID(id int) (*Notification, error)
	CreateNotification(notification *CreateNotificationPayload) error
	UpdateNotification(notification *UpdateNotificationPayload) error
	DeleteNotification(id int) error
}

type Notification struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	Message   string       `json:"message"`
	Status    string       `json:"status"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CreateNotificationPayload struct {
	UserID  int    `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
	Status  string `json:"status" validate:"required"`
}

type UpdateNotificationPayload struct {
	ID      int    `json:"id" validate:"required"`
	UserID  int    `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
	Status  string `json:"status" validate:"required"`
}

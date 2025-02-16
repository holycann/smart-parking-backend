package notifications

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

func scanRowIntoNotification(row *sql.Rows) (*Notification, error) {
	notification := new(Notification)

	err := row.Scan(
		&notification.ID,
		&notification.UserID,
		&notification.Message,
		&notification.Status,
		&notification.CreatedAt,
		&notification.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (s *Store) GetAllNotification() (notification []*Notification, err error) {
	rows, err := s.db.Query("SELECT * FROM notifications")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		sp, err := scanRowIntoNotification(rows)
		if err != nil {
			return nil, err
		}
		notification = append(notification, sp)
	}

	return notification, nil
}

func (s *Store) GetNotificationByMessage(message string) (*Notification, error) {
	rows, err := s.db.Query("SELECT * FROM notifications WHERE message = $1", message)
	if err != nil {
		return nil, err
	}

	n := new(Notification)
	for rows.Next() {
		n, err = scanRowIntoNotification(rows)
		if err != nil {
			return nil, err
		}
	}

	if n.Status == "" {
		return nil, fmt.Errorf("Notification Status not found")
	}

	return n, nil
}

func (s *Store) GetNotificationByID(id int) (*Notification, error) {
	rows, err := s.db.Query("SELECT * FROM notifications WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	sp := new(Notification)
	for rows.Next() {
		sp, err = scanRowIntoNotification(rows)
		if err != nil {
			return nil, err
		}
	}

	if sp.ID == 0 {
		return nil, fmt.Errorf("Notification not found")
	}

	return sp, nil
}

func (s *Store) CreateNotification(notification *CreateNotificationPayload) error {
	_, err := s.db.Exec("INSERT INTO notifications (user_id, message, status) VALUES ($1, $2, $3)", notification.UserID, notification.Message, notification.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateNotification(notification *UpdateNotificationPayload) error {
	_, err := s.db.Exec("UPDATE notifications SET user_id = $1, message = $2, status = $3 WHERE id = $4", notification.UserID, notification.Message, notification.Status, notification.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteNotification(id int) error {
	_, err := s.db.Exec("DELETE FROM notifications WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

package reservations

import (
	"database/sql"
)

type ReservationRepositoryInterface interface {
	GetAllReservation() ([]*Reservation, error)
	GetReservationByStartTime(StartTime int64) (*Reservation, error)
	GetReservationByID(id int) (*Reservation, error)
	CreateReservation(reservation *CreateReservationPayload) error
	UpdateReservation(reservation *UpdateReservationPayload) error
	DeleteReservation(id int) error
}

type ReservationServiceInterface interface {
	GetAllReservation() ([]*Reservation, error)
	GetReservationByID(id int) (*Reservation, error)
	CreateReservation(reservation *CreateReservationPayload) (string, error)
	UpdateReservation(reservation *UpdateReservationPayload) (string, error)
	DeleteReservation(id int) (string, error)
}

type Reservation struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id"`
	SpotID    int          `json:"spot_id"`
	VehicleID int          `json:"vehicle_id"`
	StartTime int64        `json:"start_time"`
	EndTime   int64        `json:"end_time"`
	Status    string       `json:"status"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CreateReservationPayload struct {
	UserID    int    `json:"user_id" validate:"required"`
	SpotID    int    `json:"spot_id" validate:"required"`
	VehicleID int    `json:"vehicle_id" validate:"required"`
	StartTime int64  `json:"start_time" validate:"required"`
	EndTime   int64  `json:"end_time" validate:"required"`
	Status    string `json:"status" validate:"required"`
}

type UpdateReservationPayload struct {
	ID        int    `json:"id" validate:"required"`
	UserID    int    `json:"user_id" validate:"required"`
	SpotID    int    `json:"spot_id" validate:"required"`
	VehicleID int    `json:"vehicle_id" validate:"required"`
	StartTime int64  `json:"start_time" validate:"required"`
	EndTime   int64  `json:"end_time" validate:"required"`
	Status    string `json:"status" validate:"required"`
}

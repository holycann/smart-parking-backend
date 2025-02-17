package reservations

import (
	"database/sql"
	"fmt"
	"time"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{
		db: db,
	}
}

func scanRowIntoReservation(row *sql.Rows) (*Reservation, error) {
	reservation := new(Reservation)

	err := row.Scan(
		&reservation.ID,
		&reservation.UserID,
		&reservation.SpotID,
		&reservation.VehicleID,
		&reservation.StartTime,
		&reservation.EndTime,
		&reservation.Status,
		&reservation.CreatedAt,
		&reservation.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *ReservationRepository) GetAllReservation() (reservation []*Reservation, err error) {
	rows, err := s.db.Query("SELECT * FROM reservations")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		sp, err := scanRowIntoReservation(rows)
		if err != nil {
			return nil, err
		}
		reservation = append(reservation, sp)
	}

	return reservation, nil
}

func (s *ReservationRepository) GetReservationByStartTime(StartTime int64) (*Reservation, error) {
	rows, err := s.db.Query("SELECT * FROM reservations WHERE start_time = $1", StartTime)
	if err != nil {
		return nil, err
	}

	r := new(Reservation)
	for rows.Next() {
		r, err = scanRowIntoReservation(rows)
		if err != nil {
			return nil, err
		}
	}

	if r.StartTime >= time.Now().Unix() {
		return nil, fmt.Errorf("Reservation Start Time not found")
	}

	return r, nil
}

func (s *ReservationRepository) GetReservationByID(id int) (*Reservation, error) {
	rows, err := s.db.Query("SELECT * FROM reservations WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	r := new(Reservation)
	for rows.Next() {
		r, err = scanRowIntoReservation(rows)
		if err != nil {
			return nil, err
		}
	}

	if r.ID == 0 {
		return nil, fmt.Errorf("Reservation not found")
	}

	return r, nil
}

func (s *ReservationRepository) CreateReservation(reservation *CreateReservationPayload) error {
	_, err := s.db.Exec("INSERT INTO reservations (user_id, spot_id, vehicle_id, start_time, end_time, status) VALUES ($1, $2, $3, $4, $5, $6)", reservation.UserID, reservation.SpotID, reservation.VehicleID, reservation.StartTime, reservation.EndTime, reservation.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReservationRepository) UpdateReservation(reservation *UpdateReservationPayload) error {
	_, err := s.db.Exec("UPDATE reservations SET user_id = $1, spot_id = $2, vehicle_id = $3, start_time = $4, end_time = $5, status = $6 WHERE id = $7", reservation.UserID, reservation.SpotID, reservation.VehicleID, reservation.StartTime, reservation.EndTime, reservation.Status, reservation.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReservationRepository) DeleteReservation(id int) error {
	_, err := s.db.Exec("DELETE FROM reservations WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

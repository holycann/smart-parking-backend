package reservations

import (
	"fmt"
	"time"
)

type ReservationService struct {
	repository ReservationRepositoryInterface
}

func NewService(repository ReservationRepositoryInterface) *ReservationService {
	return &ReservationService{repository: repository}
}

func (s *ReservationService) GetAllReservation() ([]*Reservation, error) {
	return s.repository.GetAllReservation()
}

func (s *ReservationService) GetReservationByID(id int) (*Reservation, error) {
	reservation, err := s.repository.GetReservationByID(id)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("error getting reservation by id: %v\n", err)
	}

	return reservation, nil
}

func (s *ReservationService) CreateReservation(payload *CreateReservationPayload) (string, error) {
	_, err := s.repository.GetReservationByStartTime(payload.StartTime)
	if err == nil {
		return "", fmt.Errorf("Reservation Start Time %s already exists", payload.StartTime)
	}

	err = s.repository.CreateReservation(&CreateReservationPayload{
		UserID:    payload.UserID,
		SpotID:    payload.SpotID,
		VehicleID: payload.VehicleID,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
		Status:    payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error create reservation: %v\n", err)
	}

	return fmt.Sprintf("Create reservation %s successfully", payload.StartTime), nil
}

func (s *ReservationService) UpdateReservation(payload *UpdateReservationPayload) (string, error) {
	re, err := s.repository.GetReservationByID(payload.ID)
	if err != nil {
		return "", fmt.Errorf("error get reservation by id: %v\n", err)
	}

	if re == nil {
		return "", fmt.Errorf("Reservation with ID %d does not exist", payload.ID)
	}

	if payload.StartTime <= time.Now().Unix() && payload.EndTime >= payload.StartTime {
		return "", fmt.Errorf("Reservation Start Time Must Be <= Time Now And End Time Must Be >= Start Time")
	}

	err = s.repository.UpdateReservation(&UpdateReservationPayload{
		ID:        payload.ID,
		UserID:    payload.UserID,
		SpotID:    payload.SpotID,
		VehicleID: payload.VehicleID,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
		Status:    payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error update reservation: %v\n", err)
	}

	return fmt.Sprintf("Update reservation %v successfully", re.StartTime), nil
}

func (s *ReservationService) DeleteReservation(id int) (string, error) {
	err := s.repository.DeleteReservation(id)
	if err != nil {
		return "", fmt.Errorf("error delete reservation: %v\n", err)
	}

	return fmt.Sprintf("Delete reservation successfully"), nil
}

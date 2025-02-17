package spots

import (
	"fmt"
)

type SpotService struct {
	repository SpotRepositoryInterface
}

func NewService(repository SpotRepositoryInterface) *SpotService {
	return &SpotService{repository: repository}
}

func (s *SpotService) GetAllSpot() ([]*Spot, error) {
	return s.repository.GetAllSpot()
}

func (s *SpotService) GetSpotByID(id int) (*Spot, error) {
	spot, err := s.repository.GetSpotByID(id)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("error getting spot by id: %v\n", err)
	}

	return spot, err
}

func (s *SpotService) CreateSpot(payload *CreateSpotPayload) (string, error) {
	_, err := s.repository.GetSpotByNumber(payload.SpotNumber)
	if err == nil {
		return "", fmt.Errorf("Spot Number %s already exists", payload.SpotNumber)
	}

	err = s.repository.CreateSpot(&CreateSpotPayload{
		ZoneID:     payload.ZoneID,
		SpotNumber: payload.SpotNumber,
		Status:     payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error create spot: %v\n", err)
	}

	return fmt.Sprintf("Create spot %s successfully", payload.SpotNumber), nil
}

func (s *SpotService) UpdateSpot(payload *UpdateSpotPayload) (string, error) {
	v, err := s.repository.GetSpotByID(payload.ID)
	if err != nil {
		return "", fmt.Errorf("error get spot by id: %v\n", err)
	}

	if v == nil {
		return "", fmt.Errorf("Spot with ID %d does not exist", payload.ID)
	}

	if payload.ZoneID == 0 && payload.SpotNumber == "" {
		return "", fmt.Errorf("Spot Zone And Number Cannot Be Empty!")
	}

	err = s.repository.UpdateSpot(&UpdateSpotPayload{
		ID:         payload.ID,
		ZoneID:     payload.ZoneID,
		SpotNumber: payload.SpotNumber,
		Status:     payload.Status,
	})
	if err != nil {
		return "", fmt.Errorf("error update spot: %v\n", err)
	}

	return fmt.Sprintf("Update spot %s successfully", v.SpotNumber), nil
}

func (s *SpotService) DeleteSpot(id int) (string, error) {
	err := s.repository.DeleteSpot(id)
	if err != nil {
		return "", fmt.Errorf("error delete spot: %v\n", err)
	}

	return fmt.Sprintf("Delete spot successfully"), nil
}

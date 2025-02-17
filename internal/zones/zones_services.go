package zones

import (
	"fmt"
)

type ZoneService struct {
	repository ZoneRepositoryInterface
}

func NewService(repository ZoneRepositoryInterface) *ZoneService {
	return &ZoneService{repository: repository}
}

func (s *ZoneService) GetAllZone() ([]*Zone, error) {
	return s.repository.GetAllZone()
}

func (s *ZoneService) GetZoneByID(id int) (*Zone, error) {
	zone, err := s.repository.GetZoneByID(id)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("error getting zone by id: %v\n", err)
	}

	return zone, nil
}

func (s *ZoneService) CreateZone(payload *CreateZonePayload) (string, error) {
	_, err := s.repository.GetZoneByName(payload.Name)
	if err == nil {
		return "", fmt.Errorf("Zone Name %s already exists", payload.Name)
	}

	err = s.repository.CreateZone(&CreateZonePayload{
		Name:       payload.Name,
		Location:   payload.Location,
		TotalSpots: payload.TotalSpots,
	})
	if err != nil {
		return "", fmt.Errorf("error create user: %v\n", err)
	}

	return fmt.Sprintf("Create zone %s successfully", payload.Name), nil
}

func (s *ZoneService) UpdateZone(payload *UpdateZonePayload) (string, error) {
	z, err := s.repository.GetZoneByID(payload.ID)
	if err != nil {
		return "", fmt.Errorf("error get zone by id: %v\n", err)
	}

	if z == nil {
		return "", fmt.Errorf("Zone wits ID %d does not exist", payload.ID)
	}

	if payload.Name == "" && payload.Location == "" {
		return "", fmt.Errorf("Zone Name And Location Cannot Be Empty!")
	}

	err = s.repository.UpdateZone(&UpdateZonePayload{
		ID:         payload.ID,
		Name:       payload.Name,
		Location:   payload.Location,
		TotalSpots: payload.TotalSpots,
	})
	if err != nil {
		return "", fmt.Errorf("error update zone: %v\n", err)
	}

	return fmt.Sprintf("Update zone %s successfully", z.Name), nil
}

func (s *ZoneService) DeleteZone(id int) (string, error) {
	err := s.repository.DeleteZone(id)
	if err != nil {
		return "", fmt.Errorf("error delete zone: %v\n", err)
	}

	return fmt.Sprintf("Delete zone successfully"), nil
}

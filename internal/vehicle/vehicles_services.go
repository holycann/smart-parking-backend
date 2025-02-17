package vehicles

import (
	"fmt"
)

type VehicleService struct {
	repository VehicleRepositoryInterface
}

func NewService(repository VehicleRepositoryInterface) *VehicleService {
	return &VehicleService{repository: repository}
}

func (s *VehicleService) GetAllVehicle() ([]*Vehicle, error) {
	return s.repository.GetAllVehicle()
}

func (s *VehicleService) GetVehicleByID(id int) (*Vehicle, error) {
	vehicle, err := s.repository.GetVehicleByID(id)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("error getting vehicle by id: %v\n", err)
	}

	return vehicle, nil
}

func (s *VehicleService) CreateVehicle(payload *CreateVehiclePayload) (string, error) {
	_, err := s.repository.GetVehicleByPlate(payload.PlateNumber)
	if err == nil {
		return "", fmt.Errorf("Vehicle Plate %s already exists", payload.PlateNumber)
	}

	err = s.repository.CreateVehicle(&CreateVehiclePayload{
		UserID:      payload.UserID,
		PlateNumber: payload.PlateNumber,
		Type:        payload.Type,
		Brand:       payload.Brand,
		Model:       payload.Model,
		Color:       payload.Color,
	})
	if err != nil {
		return "", fmt.Errorf("error create vehicle: %v\n", err)
	}

	return fmt.Sprintf("Create vehicle %s successfully", payload.PlateNumber), nil
}

func (s *VehicleService) UpdateVehicle(payload *UpdateVehiclePayload) (string, error) {
	v, err := s.repository.GetVehicleByID(payload.ID)
	if err != nil {
		return "", fmt.Errorf("error get vehicle by id: %v\n", err)
	}

	if v == nil {
		return "", fmt.Errorf("Vehicle wits ID %d does not exist", payload.ID)
	}

	if payload.UserID == 0 && payload.PlateNumber == "" {
		return "", fmt.Errorf("Vehicle User And Plat Number Cannot Be Empty!")
	}

	err = s.repository.UpdateVehicle(&UpdateVehiclePayload{
		ID:          payload.ID,
		UserID:      payload.UserID,
		PlateNumber: payload.PlateNumber,
		Type:        payload.Type,
		Brand:       payload.Brand,
		Model:       payload.Model,
		Color:       payload.Color,
	})
	if err != nil {
		return "", fmt.Errorf("error update vehicle: %v\n", err)
	}

	return fmt.Sprintf("Update vehicle %s successfully", v.PlateNumber), nil
}

func (s *VehicleService) DeleteVehicle(id int) (string, error) {
	err := s.repository.DeleteVehicle(id)
	if err != nil {
		return "", fmt.Errorf("error delete vehicle: %v\n", err)
	}

	return fmt.Sprintf("Delete vehicle successfully"), nil
}

package vehicles

import (
	"database/sql"
	"fmt"
)

type VehicleRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *VehicleRepository {
	return &VehicleRepository{
		db: db,
	}
}

func scanRowIntoVehicle(row *sql.Rows) (*Vehicle, error) {
	vehicle := new(Vehicle)

	err := row.Scan(
		&vehicle.ID,
		&vehicle.UserID,
		&vehicle.PlateNumber,
		&vehicle.Type,
		&vehicle.Brand,
		&vehicle.Model,
		&vehicle.Color,
		&vehicle.CreatedAt,
		&vehicle.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (s *VehicleRepository) GetAllVehicle() (vehicles []*Vehicle, err error) {
	rows, err := s.db.Query("SELECT * FROM vehicles")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		v, err := scanRowIntoVehicle(rows)
		if err != nil {
			return nil, err
		}
		vehicles = append(vehicles, v)
	}

	return vehicles, nil
}

func (s *VehicleRepository) GetVehicleByPlate(plate string) (*Vehicle, error) {
	rows, err := s.db.Query("SELECT * FROM vehicles WHERE plate_number = $1", plate)
	if err != nil {
		return nil, err
	}

	v := new(Vehicle)
	for rows.Next() {
		v, err = scanRowIntoVehicle(rows)
		if err != nil {
			return nil, err
		}
	}

	if v.PlateNumber == "" {
		return nil, fmt.Errorf("Vehicle plate number not found")
	}

	return v, nil
}

func (s *VehicleRepository) GetVehicleByID(id int) (*Vehicle, error) {
	rows, err := s.db.Query("SELECT * FROM vehicles WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	v := new(Vehicle)
	for rows.Next() {
		v, err = scanRowIntoVehicle(rows)
		if err != nil {
			return nil, err
		}
	}

	if v.ID == 0 {
		return nil, fmt.Errorf("Vehicle not found")
	}

	return v, nil
}

func (s *VehicleRepository) CreateVehicle(vehicle *CreateVehiclePayload) error {
	_, err := s.db.Exec("INSERT INTO vehicles (user_id, plate_number, type, brand, model, color) VALUES ($1, $2, $3, $4, $5, $6)", vehicle.UserID, vehicle.PlateNumber, vehicle.Type, vehicle.Brand, vehicle.Model, vehicle.Color)
	if err != nil {
		return err
	}

	return nil
}

func (s *VehicleRepository) UpdateVehicle(vehicle *UpdateVehiclePayload) error {
	_, err := s.db.Exec("UPDATE vehicles SET user_id = $1, plate_number = $2, type = $3 , brand = $4 , model = $5 , color = $6 WHERE id = $7", vehicle.UserID, vehicle.PlateNumber, vehicle.Type, vehicle.Brand, vehicle.Brand, vehicle.Model, vehicle.Color, vehicle.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *VehicleRepository) DeleteVehicle(id int) error {
	_, err := s.db.Exec("DELETE FROM vehicles WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

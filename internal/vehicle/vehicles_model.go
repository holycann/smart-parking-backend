package vehicles

import "database/sql"

type VehicleRepositoryInterface interface {
	GetAllVehicle() ([]*Vehicle, error)
	GetVehicleByPlate(plate string) (*Vehicle, error)
	GetVehicleByID(id int) (*Vehicle, error)
	CreateVehicle(vehicle *CreateVehiclePayload) error
	UpdateVehicle(vehicle *UpdateVehiclePayload) error
	DeleteVehicle(id int) error
}

type VehicleServiceInterface interface {
	GetAllVehicle() ([]*Vehicle, error)
	GetVehicleByID(id int) (*Vehicle, error)
	CreateVehicle(vehicle *CreateVehiclePayload) (string, error)
	UpdateVehicle(vehicle *UpdateVehiclePayload) (string, error)
	DeleteVehicle(id int) (string, error)
}

type Vehicle struct {
	ID          int          `json:"id"`
	UserID      int          `json:"user_id"`
	PlateNumber string       `json:"plat_number"`
	Type        string       `json:"type"`
	Brand       string       `json:"brand"`
	Model       string       `json:"model"`
	Color       string       `json:"color"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type CreateVehiclePayload struct {
	UserID      int    `json:"user_id" validate:"required"`
	PlateNumber string `json:"plate_number" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Model       string `json:"model" validate:"required"`
	Color       string `json:"color" validate:"required"`
}

type UpdateVehiclePayload struct {
	ID          int    `json:"id" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
	PlateNumber string `json:"plate_number" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Model       string `json:"model" validate:"required"`
	Color       string `json:"color" validate:"required"`
}

package zones

import "database/sql"

type ZoneRepositoryInterface interface {
	GetAllZone() ([]*Zone, error)
	GetZoneByName(name string) (*Zone, error)
	GetZoneByID(id int) (*Zone, error)
	CreateZone(zone *CreateZonePayload) error
	UpdateZone(zone *UpdateZonePayload) error
	DeleteZone(id int) error
}

type ZoneServiceInterface interface {
	GetAllZone() ([]*Zone, error)
	GetZoneByID(id int) (*Zone, error)
	CreateZone(zone *CreateZonePayload) (string, error)
	UpdateZone(zone *UpdateZonePayload) (string, error)
	DeleteZone(id int) (string, error)
}

type Zone struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Location   string       `json:"location"`
	TotalSpots int          `json:"total_spots"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type CreateZonePayload struct {
	Name       string `json:"name" validate:"required"`
	Location   string `json:"location" validate:"required"`
	TotalSpots int    `json:"total_spots" validate:"required,number"`
}

type UpdateZonePayload struct {
	ID         int    `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Location   string `json:"location" validate:"required"`
	TotalSpots int    `json:"total_spots" validate:"required,number"`
}

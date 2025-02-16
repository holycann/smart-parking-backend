package spots

import "database/sql"

type SpotStore interface {
	GetAllSpot() ([]*Spot, error)
	GetSpotByNumber(plate string) (*Spot, error)
	GetSpotByID(id int) (*Spot, error)
	CreateSpot(spot *CreateSpotPayload) error
	UpdateSpot(spot *UpdateSpotPayload) error
	DeleteSpot(id int) error
}

type Spot struct {
	ID         int          `json:"id"`
	ZoneID     int          `json:"zone_id"`
	SpotNumber string       `json:"spot_number"`
	Status     string       `json:"status"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type CreateSpotPayload struct {
	ZoneID     int    `json:"zone_id" validate:"required"`
	SpotNumber string `json:"spot_number" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

type UpdateSpotPayload struct {
	ID         int    `json:"id" validate:"required"`
	ZoneID     int    `json:"zone_id" validate:"required"`
	SpotNumber string `json:"spot_number" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

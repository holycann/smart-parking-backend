package spots

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func scanRowIntoSpot(row *sql.Rows) (*Spot, error) {
	spot := new(Spot)

	err := row.Scan(
		&spot.ID,
		&spot.ZoneID,
		&spot.SpotNumber,
		&spot.Status,
		&spot.CreatedAt,
		&spot.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return spot, nil
}

func (s *Store) GetAllSpot() (spot []*Spot, err error) {
	rows, err := s.db.Query("SELECT * FROM spots")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		sp, err := scanRowIntoSpot(rows)
		if err != nil {
			return nil, err
		}
		spot = append(spot, sp)
	}

	return spot, nil
}

func (s *Store) GetSpotByNumber(plate string) (*Spot, error) {
	rows, err := s.db.Query("SELECT * FROM spots WHERE spot_number = $1", plate)
	if err != nil {
		return nil, err
	}

	sp := new(Spot)
	for rows.Next() {
		sp, err = scanRowIntoSpot(rows)
		if err != nil {
			return nil, err
		}
	}

	if sp.SpotNumber == "" {
		return nil, fmt.Errorf("Spot plate number not found")
	}

	return sp, nil
}

func (s *Store) GetSpotByID(id int) (*Spot, error) {
	rows, err := s.db.Query("SELECT * FROM spots WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	sp := new(Spot)
	for rows.Next() {
		sp, err = scanRowIntoSpot(rows)
		if err != nil {
			return nil, err
		}
	}

	if sp.ID == 0 {
		return nil, fmt.Errorf("Spot not found")
	}

	return sp, nil
}

func (s *Store) CreateSpot(spot *CreateSpotPayload) error {
	_, err := s.db.Exec("INSERT INTO spots (zone_id, spot_number, status) VALUES ($1, $2, $3)", spot.ZoneID, spot.SpotNumber, spot.Status)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateSpot(spot *UpdateSpotPayload) error {
	_, err := s.db.Exec("UPDATE spots SET zone_id = $1, spot_number = $2, status = $3 WHERE id = $4", spot.ZoneID, spot.SpotNumber, spot.Status, spot.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteSpot(id int) error {
	_, err := s.db.Exec("DELETE FROM spots WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

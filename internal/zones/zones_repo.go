package zones

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

func scanRowIntoZone(row *sql.Rows) (*Zone, error) {
	zone := new(Zone)

	err := row.Scan(
		&zone.ID,
		&zone.Name,
		&zone.TotalSpots,
		&zone.CreatedAt,
		&zone.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return zone, nil
}

func (s *Store) GetAllZone() (zones []*Zone, err error) {
	rows, err := s.db.Query("SELECT * FROM zones")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		z, err := scanRowIntoZone(rows)
		if err != nil {
			return nil, err
		}
		zones = append(zones, z)
	}

	return zones, nil
}

func (s *Store) GetZoneByName(name string) (*Zone, error) {
	rows, err := s.db.Query("SELECT * FROM zones WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	z := new(Zone)
	for rows.Next() {
		z, err = scanRowIntoZone(rows)
		if err != nil {
			return nil, err
		}
	}

	if z.Name == "" {
		return nil, fmt.Errorf("Zone not found")
	}

	return z, nil
}

func (s *Store) GetZoneByID(id int) (*Zone, error) {
	rows, err := s.db.Query("SELECT * FROM zones WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	z := new(Zone)
	for rows.Next() {
		z, err = scanRowIntoZone(rows)
		if err != nil {
			return nil, err
		}
	}

	if z.ID == 0 {
		return nil, fmt.Errorf("Zone not found")
	}

	return z, nil
}

func (s *Store) CreateZone(zone *CreateZonePayload) error {
	_, err := s.db.Exec("INSERT INTO zones (name, location, total_spots) VALUES ($1, $2, $3)", zone.Name, zone.Location, zone.TotalSpots)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateZone(zone *UpdateZonePayload) error {
	_, err := s.db.Exec("UPDATE zones SET name = $1, location = $2, total_spots = $3 WHERE id = $4", zone.Name, zone.Location, zone.TotalSpots, zone.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteZone(id int) error {
	_, err := s.db.Exec("DELETE FROM zones WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

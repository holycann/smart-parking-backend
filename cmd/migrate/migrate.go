package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	sqlMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/holycann/smart-parking-backend/config"
	"github.com/holycann/smart-parking-backend/db"
)

func main() {
	// Gunakan PostgreSQL storage
	db, err := db.NewPostgresStorage(config.Env.DBAddress, config.Env.MaxOpenConns, config.Env.MaxIdleConns, config.Env.MaxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	// Gunakan driver PostgreSQL
	driver, err := sqlMigrate.WithInstance(db, &sqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Gunakan file source untuk migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations", // Lokasi folder migration
		"postgres",             // Database yang digunakan adalah PostgreSQL
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Ambil perintah dari argumen
	cmd := os.Args[len(os.Args)-1]

	// Eksekusi perintah 'up' atau 'down' untuk migrasi
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}

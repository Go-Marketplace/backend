package app

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

const (
	_defaultAttempts = 5
	_defaultTimeout  = time.Second
)

func init() {
	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		log.Fatalf("migrate: environment variable not declared: PG_URL")
	}

	var (
		attempts = _defaultAttempts
		err      error
		db       *sql.DB
	)

	for attempts > 0 {
		db, err = sql.Open("postgres", databaseURL)
		if err == nil {
			break
		}
		log.Printf("migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()

	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}

	migrationsDir := filepath.Join(dir, "migrations")

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("the migrations up attempt was successful")
}

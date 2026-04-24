package internal

import (
	"context"
	"database/sql"
	"time"

	sqlite3 "github.com/mattn/go-sqlite3"
)

func OpenDatabase(ctx context.Context, dsn string) (*sql.DB, error) {
	sqlite3.SQLiteTimestampFormats = []string{
		time.RFC3339Nano,
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	timeoutContext, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := db.PingContext(timeoutContext); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}

package service

import (
	"context"
	"database/sql"
	"time"
)

func DbGetUpdates(db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select get_updates()`

	if err := db.QueryRowContext(ctx, query).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func DbCreateUpdate(version int, url string, force bool, checksum string, changelog string, db *sql.DB) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select create_update($1, $2, $3, $4, $5)`

	if err := db.QueryRowContext(ctx, query, version, url, force, checksum, changelog).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

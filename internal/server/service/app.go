package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

func DbGetUpdates(db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select get_update()`

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

func DbGetConfig(db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select get_config()`

	if err := db.QueryRowContext(ctx, query).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func DbCreateConfig(config_group string, value json.RawMessage, db *sql.DB) (k string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select create_config($1, $2)`

	if err := db.QueryRowContext(ctx, query, config_group, value).Scan(&config_group); err != nil {
		return config_group, err
	}

	return config_group, nil
}
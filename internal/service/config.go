package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

func DbCreateConfig(config_group string, value json.RawMessage, db *sql.DB) (k string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select create_config($1, $2)`

	if err := db.QueryRowContext(ctx, query, config_group, value).Scan(&config_group); err != nil {
		return config_group, err
	}

	return config_group, nil
}

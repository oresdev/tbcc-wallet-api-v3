package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

func DbCreateVpnKey(key string, validity int, used bool, user_id uuid.UUID, txhash string, with_pro null.Bool, db *sql.DB) (k string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select create_vpn_key($1, $2, $3, $4, $5, $6)`

	if err := db.QueryRowContext(ctx, query, key, validity, used, user_id, txhash, with_pro).Scan(&key); err != nil {
		return key, err
	}

	return key, nil
}

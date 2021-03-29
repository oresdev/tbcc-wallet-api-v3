package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func DbGetAllUsers(db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select get_users()`

	if err := db.QueryRowContext(ctx, query).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func DbGetUserByID(id string, db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select get_user_by_id($1)`

	if err := db.QueryRowContext(ctx, query, id).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func DbGetUserExt(id string, db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select get_user_ext($1)`

	if err := db.QueryRowContext(ctx, query, id).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func DbCreateUser(useraddress []string, accounttype string, smartcard bool, db *sql.DB) (id uuid.UUID, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select create_user($1, $2, $3)`

	if err := db.QueryRowContext(ctx, query, pq.Array(useraddress), accounttype, smartcard).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

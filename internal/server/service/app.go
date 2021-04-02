package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

// DbGetUpdates ...
func DbGetUpdates(db *sql.DB) (data []byte, err error) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	query := `select app_update_get_rows()`
	err = db.QueryRow(query).Scan(&data)

	if(err != nil){
		return nil, err
	}

	//if err := db.QueryRowContext(ctx, query).Scan(&data); err != nil {
	//	return nil, err
	//}

	return data, nil
}

// DbCreateUpdate ...
func DbCreateUpdate(version int, url string, force bool, checksum string, changelog string, db *sql.DB) (id int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := `select app_update_create_row($1, $2, $3, $4, $5)`

	if err := db.QueryRowContext(ctx, query, version, url, force, checksum, changelog).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

// DbGetConfig ...
func DbGetConfig(db *sql.DB) (data []byte, err error) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	query := `select app_config_get_rows()`
	err = db.QueryRow(query).Scan(&data)
	if(err != nil){
		return nil, err
	}

	//if err := db.QueryRowContext(ctx, query).Scan(&data); err != nil {
	//	return nil, err
	//}

	return data, nil
}

// DbCreateConfig ...
func DbCreateConfig(config_group string, value json.RawMessage, db *sql.DB) (k string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := `select app_config_create_row($1, $2)`

	if err := db.QueryRowContext(ctx, query, config_group, value).Scan(&config_group); err != nil {
		return config_group, err
	}

	return config_group, nil
}

// DbCountVersion ...
func DbCountVersion(version int, db *sql.DB) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	query := `select app_counter_update_row($1)`

	db.QueryRow(query)
	

//	db.QueryRowContext(ctx, query, version)

	return
}

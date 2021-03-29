package service

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/model"
	"gopkg.in/guregu/null.v4"
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

func DbGetUserByAddress(address string, db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `select get_user_by_address($1)`

	if err := db.QueryRowContext(ctx, query, address).Scan(&data); err != nil {
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

func DbMigrateUser(addresses []string, db *sql.DB) (data []byte, err error) {

	users := []model.UserMigrate{}

	conditionArr := make([]string, len(addresses))

	for i, addr := range addresses {
		conditionArr[i] = " address='" + addr + "' "
	}

	clients, err := db.Query("select id, address, paid_fee, paid_smart_card from clients where" + strings.Join(conditionArr, "OR"))

	for clients.Next() {
		user := model.UserMigrate{}
		err := clients.Scan(&user.ID, &user.Address, &user.PaidFee, &user.PaidSmartcard)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	newUser := model.User{}

	newUser.Useraddress = addresses

	for _, u := range users {
		if u.PaidFee != null.FloatFromPtr(nil) {
			if u.PaidFee == null.FloatFrom(1.5) && newUser.Accounttype != "PRO" {
				newUser.Accounttype = "Premium"
			}
			if u.PaidFee == null.FloatFrom(2) {
				newUser.Accounttype = "PRO"
			}
		}
		if u.PaidSmartcard != null.FloatFromPtr(nil) {
			newUser.Smartcard = true
		}
	}

	uuid, err := DbCreateUser(newUser.Useraddress, newUser.Accounttype, newUser.Smartcard, db)

	data, err = DbGetUserExt(uuid.String(), db)

	return data, nil
}

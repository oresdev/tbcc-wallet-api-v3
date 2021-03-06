package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/model"
	"gopkg.in/guregu/null.v4"
)

// DbGetAllUsers ...
func DbGetAllUsers(db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := `select users_get_rows()`

	if err := db.QueryRowContext(ctx, query).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbGetUserByID ...
func DbGetUserByID(id string, db *sql.DB) (data []byte, err error) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	query := `select users_get_by_uuid($1)`
	err = db.QueryRow(query, id).Scan(&data)
	if(err != nil){
		return nil, err
	}
	//if err := db.QueryRowContext(ctx, query, id).Scan(&data); err != nil {
//		return nil, err
//	}

	return data, nil
}

// DbGetUserExt ...
func DbGetUserExt(uuid string, db *sql.DB) (data []byte, err error) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	//defer cancel()

	query := `select users_get_extended_by_uuid($1)`
	err = db.QueryRow(query, uuid).Scan(&data)

	if(err != nil){
		return nil, err
	}
	//if err := db.QueryRowContext(ctx, query, uuid).Scan(&data); err != nil {
//		return nil, err
//	}

	return data, nil
}

// DbUpdateUser ...
func DbUpdateUser(uuid string, address string, db *sql.DB) (data []byte, err error) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	query := `select users_update_by_uuid($1, $2)`

	err = db.QueryRow(query, uuid,address).Scan(&data)
	if(err!= nil){
		return nil,err
	}
	//if err := db.QueryRowContext(ctx, query, uuid, address).Scan(&data); err != nil {
//		return nil, err
//	}

	return data, nil
}

// DbCreateUser ...
func DbCreateUser(useraddress []string, accounttype string, smartcard bool, db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := `select users_create_row($1, $2, $3)`

	if err := db.QueryRowContext(ctx, query, pq.Array(useraddress), accounttype, smartcard).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbMigrateUser ...
func DbMigrateUser(addresses []string, db *sql.DB) (data []byte, err error) {

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	//defer cancel()

	


	query := "select users_check_exists_by_addresses($1)"

	err = db.QueryRow(query, pq.Array(addresses)).Scan(&data)

	if(err == nil && len(data)>0){
		return data, nil
	}
	//err = db.QueryRowContext(ctx, query, pq.Array(addresses)).Scan(&data)
	//if err == nil && len(data) > 0 {
	//	return data, nil
	//}

	users := []model.UserMigrate{}

	conditionArr := make([]string, len(addresses))

	for i, addr := range addresses {
		conditionArr[i] = " address='" + addr + "' "
	}

	clients, err := db.Query("select id, address, paid_fee, paid_smart_card from public.clients where" + strings.Join(conditionArr, "OR"))

	newUser := model.User{}

	newUser.Useraddress = addresses

	newUser.Accounttype = "Free"

	if clients != nil {

		for clients.Next() {
			user := model.UserMigrate{}

			if err := clients.Scan(&user.ID, &user.Address, &user.PaidFee, &user.PaidSmartcard); err != nil {
				return nil, err
			}

			users = append(users, user)
		}

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

		data, err = DbCreateUser(newUser.Useraddress, newUser.Accounttype, newUser.Smartcard, db)

		json.Unmarshal(data, &newUser)

		if len(users) > 0 {

			conditionArrVpn := make([]string, len(users))

			for i, u := range users {
				conditionArrVpn[i] = " user_id=" + strconv.Itoa(u.ID) + " "
			}

			vpn_keys, _ := db.Query("select id from public.vpn_keys where" + strings.Join(conditionArrVpn, "OR"))

			if vpn_keys != nil {

				keys := []int{}

				for vpn_keys.Next() {

					var key int

					if err := vpn_keys.Scan(&key); err != nil {
						return nil, err
					}

					keys = append(keys, key)
				}

				if len(keys) > 0 {

					conditionArrUpdateVpn := make([]string, len(keys))

					for i, k := range keys {
						conditionArrUpdateVpn[i] = " id=" + strconv.Itoa(k) + " "
					}

					_, err = db.Query("update vpn_keys set user_id = $1 where"+strings.Join(conditionArrUpdateVpn, "OR"), newUser.ID.String())

				}
			}

		}
	}

	data, err = DbGetUserExt(newUser.ID.String(), db)

	return data, nil
}

// DbUpdateVpnKey ...
func DbUpdateVpnKey(uuid string, txhash string, db *sql.DB) (data []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := `select vpn_keys_update_by_uuid($1, $2)`

	if err := db.QueryRowContext(ctx, query, uuid, txhash).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

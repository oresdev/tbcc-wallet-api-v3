package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/model"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/service"
	"github.com/sirupsen/logrus"
)

// GetUsersHandler ...
func GetUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := service.DbGetAllUsers(db)
		if err != nil {
			logrus.Errorf("GetUsersHandler db: %v", err)
			http.Error(w, "GetUsersHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(users); err != nil {
			logrus.Errorf("GetUsersHandler write: %v", err)
			http.Error(w, "GetUsersHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// GetUserHandler ...
func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "uuid")

		user, err := service.DbGetUserByID(id, db)

		if err != nil {
			logrus.Errorf("getUserHandler db: %v", err)
			http.Error(w, "getUserHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(user); err != nil {
			logrus.Errorf("GetUserHandler write: %v", err)
			http.Error(w, "GetUserHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// GetExtendedUserHandler ...
func GetExtendedUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "uuid")

		user, err := service.DbGetUserExt(id, db)

		if err != nil {
			logrus.Errorf("GetExtendedUserHandler db: %v", err)
			http.Error(w, "GetExtendedUserHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(user); err != nil {
			logrus.Errorf("GetExtendedUserHandler write: %v", err)
			http.Error(w, "GetExtendedUserHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// UpdateUserHandler ...
func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := chi.URLParam(r, "uuid")

		address := struct {
			address string
		}{}

		if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
			http.Error(w, "UpdateUserHandler read invalid params", http.StatusBadRequest)
			return
		}

		user, err := service.DbUpdateUser(uuid, address.address, db)
		if err != nil {
			logrus.Errorf("DbUpdateUser: %v", err)
			http.Error(w, "DbUpdateUser err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(user); err != nil {
			logrus.Errorf("UpdateUserHandler write uuid: %v", err)
			http.Error(w, "UpdateUserHandler write uuid", http.StatusInternalServerError)
			return
		}
	}
}

// CreateUserHandler ...
func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := model.User{}

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "CreateUserHandler read invalid params", http.StatusBadRequest)
			return
		}

		user, err := service.DbCreateUser(p.Useraddress, p.Accounttype, p.Smartcard, db)
		if err != nil {
			logrus.Errorf("createUserHandler db: %v", err)
			http.Error(w, "createUserHandler err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(user); err != nil {
			logrus.Errorf("CreateUserHandler write id: %v", err)
			http.Error(w, "CreateUserHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

// MigrateUserHandler ...
func MigrateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := model.UserMigrateBody{}

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {

			http.Error(w, "MigrateUserHandler read invalid params", http.StatusBadRequest)
			return
		}

		newUser, err := service.DbMigrateUser(p.Addresses, db)
		if err != nil {
			logrus.Errorf("MigrateUserHandler db: %v", err)
			http.Error(w, "MigrateUserHandler err", http.StatusInternalServerError)
			return
		}

		u := model.User{}

		json.Unmarshal(newUser, &u)

		tmp, err := service.DbGetUserExt(u.ID.String(), db)
		if err != nil {
			logrus.Errorf("MigrateUserHandler db: %v", err)
			http.Error(w, "MigrateUserHandler err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(tmp); err != nil {
			logrus.Errorf("MigrateUserHandler write id: %v", err)
			http.Error(w, "MigrateUserHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

// BuyVPNKeysHandler ...
func BuyVPNKeysHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := chi.URLParam(r, "uuid")

		v := model.VpnKeyBuyBody{}

		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			http.Error(w, "BuyVPNKeysHandler read invalid params", http.StatusBadRequest)
			return
		}

		key, err := service.DbUpdateVpnKey(v.TxHash, uuid, db)
		if err != nil {
			logrus.Errorf("BuyVPNKeysHandler db: %v", err)
			http.Error(w, "BuyVPNKeysHandler err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(&key); err != nil {
			logrus.Errorf("CreateConfigHandler write id: %v", err)
			http.Error(w, "CreateConfigHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

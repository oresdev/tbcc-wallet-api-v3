package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/model"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/service"
	"github.com/sirupsen/logrus"
)

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

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "uid")

		ext := r.URL.Query().Get("ext")

		var user []byte
		var err error

		if ext == "true" {
			user, err = service.DbGetUserExt(id, db)
		} else {
			user, err = service.DbGetUserByID(id, db)
		}

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

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := model.User{}

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "CreateUserHandler read invalid params", http.StatusBadRequest)
			return
		}

		id, err := service.DbCreateUser(p.Useraddress, p.Accounttype, p.Smartcard, db)
		if err != nil {
			logrus.Errorf("createUserHandler db: %v", err)
			http.Error(w, "createUserHandler err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(&id); err != nil {
			logrus.Errorf("CreateUserHandler write id: %v", err)
			http.Error(w, "CreateUserHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

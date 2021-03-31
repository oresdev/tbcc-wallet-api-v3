package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/model"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/service"
	"github.com/sirupsen/logrus"
)

// GetUpdateHandler ...
func GetUpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := service.DbGetUpdates(db)
		if err != nil {
			logrus.Errorf("GetUpdateHandler db: %v", err)
			http.Error(w, "GetUpdateHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(data); err != nil {
			logrus.Errorf("GetUpdateHandler write: %v", err)
			http.Error(w, "GetUpdateHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// CreateUpdateHandler ...
func CreateUpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := model.Updates{}

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "CreateConfigHandler read invalid params", http.StatusBadRequest)
			return
		}

		id, err := service.DbCreateUpdate(p.Version, p.URL, p.Force, p.Checksum, p.ChangeLog, db)
		if err != nil {
			logrus.Errorf("CreateConfigHandler db: %v", err)
			http.Error(w, "CreateConfigHandler err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(&id); err != nil {
			logrus.Errorf("CreateConfigHandler write id: %v", err)
			http.Error(w, "CreateConfigHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

// GetConfigHandler ...
func GetConfigHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conf, err := service.DbGetConfig(db)
		if err != nil {
			logrus.Errorf("GetConfigHandler db: %v", err)
			http.Error(w, "GetConfigHandler err", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(conf); err != nil {
			logrus.Errorf("GetConfigHandler write: %v", err)
			http.Error(w, "GetConfigHandler write err", http.StatusInternalServerError)
			return
		}
	}
}

// CreateConfigHandler ...
func CreateConfigHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := model.Configs{}

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "CreateConfigHandler read invalid params", http.StatusBadRequest)
			return
		}

		id, err := service.DbCreateConfig(p.ConfigGroup, p.Value, db)
		if err != nil {
			logrus.Errorf("CreateConfigHandler db: %v", err)
			http.Error(w, "CreateConfigHandler err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(&id); err != nil {
			logrus.Errorf("CreateConfigHandler write id: %v", err)
			http.Error(w, "CreateConfigHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

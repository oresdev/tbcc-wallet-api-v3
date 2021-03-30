package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/oresdev/tbcc-wallet-api-v3/internal/model"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/service"
	"github.com/sirupsen/logrus"
)

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

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(&id); err != nil {
			logrus.Errorf("CreateConfigHandler write id: %v", err)
			http.Error(w, "CreateConfigHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

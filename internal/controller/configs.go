package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/oresdev/tbcc-wallet-api-v3/internal/model"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/service"
	"github.com/sirupsen/logrus"
)

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

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(&id); err != nil {
			logrus.Errorf("CreateConfigHandler write id: %v", err)
			http.Error(w, "CreateConfigHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

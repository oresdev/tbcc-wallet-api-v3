package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/oresdev/tbcc-wallet-api-v3/internal/model"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/service"
	"github.com/sirupsen/logrus"
)

func CreateVpnKeyHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := model.VpnKey{}

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "CreateVpnKeyHandler read invalid params", http.StatusBadRequest)
			return
		}

		id, err := service.DbCreateVpnKey(p.Key, p.Validity, p.Used, p.UserID, p.TxHash, p.WithPro, db)
		if err != nil {
			logrus.Errorf("CreateVpnKeyHandler db: %v", err)
			http.Error(w, "CreateVpnKeyHandler err", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(&id); err != nil {
			logrus.Errorf("CreateVpnKeyHandler write id: %v", err)
			http.Error(w, "CreateVpnKeyHandler write id", http.StatusInternalServerError)
			return
		}
	}
}

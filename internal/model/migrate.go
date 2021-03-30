package model

import (
	"gopkg.in/guregu/null.v4"
)

// UserMigrateBody struct
type UserMigrateBody struct {
	Addresses []string `json:"addresses"`
}

// UserMigrate struct
type UserMigrate struct {
	ID            int
	Address       string
	PaidFee       null.Float
	PaidSmartcard null.Float
}

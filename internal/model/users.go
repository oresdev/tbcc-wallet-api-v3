package model

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

// User struct
type User struct {
	ID          uuid.UUID `json:"id" sql:",type:uuid"`
	Useraddress []string  `json:"useraddress"`
	Accounttype string    `json:"accounttype"`
	Smartcard   bool      `json:"smartcard"`
	VpnKeys     []VpnKey  `json:"vpn_keys"`
}

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

// p.ID, p.Address, p.PaidFee, p.GotTbcFee, p.PaidSmartcard, p.GotTbcSmartcard, p.FundsRestored, p.NeedCheckBrokenMnemonic, p.TbcBepSwapped, p.LotteryAccepted,

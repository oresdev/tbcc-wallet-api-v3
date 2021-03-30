package model

import (
	"github.com/google/uuid"
)

// User struct
type User struct {
	ID          uuid.UUID `json:"id" sql:",type:uuid"`
	Useraddress []string  `json:"useraddress"`
	Accounttype string    `json:"accounttype"`
	Smartcard   bool      `json:"smartcard"`
	VpnKeys     []VpnKey  `json:"vpn_keys"`
}

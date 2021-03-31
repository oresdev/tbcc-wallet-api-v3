package model

import (
	"gopkg.in/guregu/null.v4"

	"github.com/google/uuid"
)

// VpnKey struct
type VpnKey struct {
	ID        int       `json:"id"`
	Key       string    `json:"key"`
	Validity  int       `json:"validity"`
	Used      bool      `json:"used"`
	UserID    uuid.UUID `json:"user_id" sql:",type:uuid"`
	TxHash    string    `json:"txhash"`
	WithPro   null.Bool `json:"with_pro"`
	Timestamp null.Time `json:"timestamp"`
}

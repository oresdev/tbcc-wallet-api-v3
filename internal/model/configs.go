package model

import "encoding/json"

// Configs struct
type Configs struct {
	ConfigGroup string          `json:"config_group"`
	Value       json.RawMessage `json:"value"`
}

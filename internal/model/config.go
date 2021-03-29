package model

import "encoding/json"

// Config struct
type Config struct {
	ConfigGroup string          `json:"config_group"`
	Value       json.RawMessage `json:"value"`
}

package model

import "encoding/json"

// Updates struct
type Updates struct {
	Version   int    `json:"version"`
	URL       string `json:"url"`
	Force     bool   `json:"force"`
	Checksum  string `json:"checksum"`
	ChangeLog string `json:"changelog"`
}

// Configs struct
type Configs struct {
	ConfigGroup string          `json:"config_group"`
	Value       json.RawMessage `json:"value"`
}

// Counter struct
type Counter struct {
	Version int `json:"version"`
}

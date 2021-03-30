package model

// Updates struct
type Updates struct {
	Version   int    `json:"version"`
	URL       string `json:"url"`
	Force     bool   `json:"force"`
	Checksum  string `json:"checksum"`
	ChangeLog string `json:"changelog"`
}

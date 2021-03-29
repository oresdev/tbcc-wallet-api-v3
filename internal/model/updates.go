package model

// Update struct
type Update struct {
	Version   int    `json:"version"`
	URL       string `json:"url"`
	Force     bool   `json:"force"`
	Checksum  string `json:"checksum"`
	ChangeLog string `json:"changelog"`
}

package domain

type ERole struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Enabled     bool    `json:"enabled"`
}

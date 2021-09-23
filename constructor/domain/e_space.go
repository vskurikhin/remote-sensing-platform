package domain

import "github.com/google/uuid"

type ESpace struct {
	Id      int64     `json:"id"`
	Uuid    uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Base    bool      `json:"base"`
	Enabled bool      `json:"enabled"`
}

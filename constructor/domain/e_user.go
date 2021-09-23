package domain

import (
	"github.com/google/uuid"
	"time"
)

type EUser struct {
	Id            int64     `json:"id"`
	Uuid          uuid.UUID `json:"uuid"`
	SpaceId       int64     `json:"spaceId"`
	Email         string    `json:"email"`
	Login         *string   `json:"login"`
	SberId        *string   `json:"sberId"`
	Department    *string   `json:"department"`
	Name          *string   `json:"name"`
	AvatarUrl     *string   `json:"avatarUrl"`
	Enabled       bool      `json:"enabled"`
	Language      string    `json:"language"`
	LastRequestAt time.Time `json:"lastRequestAt"`
	Space         ESpace    `json:"space"`
	Online        bool      `json:"online"`
	Roles         []ERole   `json:"roles"`
}

package domain

type EScreenComplete struct {
	Index              *int    `json:"index"`
	Name               *string `json:"name"`
	ButtonText         *string `json:"buttonText"`
	MyAnswerButtonText *string `json:"myAnswerButtonText"`
	Description        *string `json:"description"`
	MyAnswerEnabled    *bool   `json:"myAnswerEnabled"`
	LinkInSameScreen   *bool   `json:"linkInSameScreen"`
	Link               *string `json:"link"`
	LinkVk             *string `json:"linkVk"`
	LinkTwitter        *string `json:"linkTwitter"`
	LinkFacebook       *string `json:"linkFacebook"`
	LinkTelegram       *string `json:"linkTelegram"`
	LinkInstagram      *string `json:"linkInstagram"`
}

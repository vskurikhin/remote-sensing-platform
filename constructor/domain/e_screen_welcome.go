package domain

type EScreenWelcome struct {
	Index       *int    `json:"index"`
	Name        *string `json:"name"`
	ButtonText  *string `json:"buttonText"`
	Description *string `json:"description"`
}

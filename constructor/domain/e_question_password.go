package domain

type EQuestionPassword struct {
	LoginPlaceholder     map[string]interface{} `json:"loginPlaceholder"`
	PasswordPlaceholder  map[string]interface{} `json:"passwordPlaceholder"`
	Password             *string                `json:"password"`
	LoginIdentifierPairs []string               `json:"loginIdentifierPairs"`
}

package domain

type EGroup struct {
	Name             map[string]interface{} `json:"name"`
	ButtonText       map[string]interface{} `json:"buttonText"`
	QuestionsShuffle *bool                  `json:"questionsShuffle"`
	Indent           *bool                  `json:"indent"`
	Children         []string               `json:"children"`
}

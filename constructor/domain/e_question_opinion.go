package domain

type EQuestionOpinion struct {
	EmotionScale       *bool   `json:"emotionScale"`
	GradientEnabled    *bool   `json:"gradientEnabled"`
	GradientColorLeft  *string `json:"gradientColorLeft"`
	GradientColorRight *string `json:"gradientColorRight"`
	OptionsValueMin    *int    `json:"optionsValueMin"`
	OptionsValueMax    *int    `json:"optionsValueMax"`
	Shape              *string `json:"shape"`
}

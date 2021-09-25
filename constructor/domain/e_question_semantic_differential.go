package domain

type EQuestionSemanticDifferential struct {
	HideZero              *bool   `json:"hideZero"`
	ShowOptionsNumeration *bool   `json:"showOptionsNumeration"`
	GradientEnabled       *bool   `json:"gradientEnabled"`
	Shape                 *string `json:"shape"`
	GradientColorLeft     *bool   `json:"gradientColorLeft"`
	GradientColorRight    *bool   `json:"gradientColorRight"`
	OptionsShuffle        *bool   `json:"optionsShuffle"`
}

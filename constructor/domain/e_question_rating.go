package domain

type EQuestionRating struct {
	OptionsValueMin            *int                   `json:"optionsValueMin"`
	OptionsValueMax            *int                   `json:"optionsValueMax"`
	ShowOptionsNumeration      *bool                  `json:"showOptionsNumeration"`
	RevertOptionsNumeration    *bool                  `json:"revertOptionsNumeration"`
	StartFromOne               *bool                  `json:"startFromOne"`
	GradientEnabled            *bool                  `json:"gradientEnabled"`
	GradientColorLeft          *string                `json:"gradientColorLeft"`
	GradientColorRight         *string                `json:"gradientColorRight"`
	Color                      *string                `json:"color"`
	Step                       *float64               `json:"step"`
	SectionCount               *int                   `json:"sectionCount"`
	OptionInitialNumberValue   *int                   `json:"optionInitialNumberValue"`
	OptionInitialInterestValue *int                   `json:"optionInitialInterestValue"`
	SubType                    *string                `json:"subType"`
	SliderViewType             *string                `json:"sliderViewType"`
	Shape                      *string                `json:"shape"`
	Scale                      map[string]interface{} `json:"scale"`
}

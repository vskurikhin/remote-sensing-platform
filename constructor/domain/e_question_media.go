package domain

type EQuestionMedia struct {
	OptionsShuffle              *bool `json:"optionsShuffle"`
	OptionsMultipleSelect       *bool `json:"optionsMultipleSelect"`
	OptionsMultipleSelectMin    *int  `json:"optionsMultipleSelectMin"`
	OptionsMultipleSelectMax    *int  `json:"optionsMultipleSelectMax"`
	OptionsIncreased            *bool `json:"optionsIncreased"`
	OptionsFullScreenEnabled    *bool `json:"optionsFullScreenEnabled"`
	OptionsLandscapeFormat      *bool `json:"optionsLandscapeFormat"`
	OptionsSelectButtonDisabled *bool `json:"optionsSelectButtonDisabled"`
}

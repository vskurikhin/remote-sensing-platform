package domain

type EQuestionClosed struct {
	OptionsIndexType          *string `json:"optionsIndexType"`
	OptionsShowType           *string `json:"optionsShowType"`
	OptionsSortAlphabetically *bool   `json:"optionsSortAlphabetically"`
	OptionsShuffle            *bool   `json:"optionsShuffle"`
	OptionsHorizontal         *bool   `json:"optionsHorizontal"`
	OptionsMultipleSelect     *bool   `json:"optionsMultipleSelect"`
	OptionsMultipleSelectMin  *int    `json:"optionsMultipleSelectMin"`
	OptionsMultipleSelectMax  *int    `json:"optionsMultipleSelectMax"`
}

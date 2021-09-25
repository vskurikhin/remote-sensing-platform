package domain

type EQuestionOpened struct {
	OptionsMultipleSelectMin *int  `json:"optionsMultipleSelectMin"`
	SequentialDisplay        *bool `json:"sequentialDisplay"`
	Multiline                *bool `json:"multiline"`
	HideTip                  *bool `json:"hideTip"`
	StringMinLength          *int  `json:"stringMinLength"`
	StringMaxLength          *int  `json:"stringMaxLength"`
}

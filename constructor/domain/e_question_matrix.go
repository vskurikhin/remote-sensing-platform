package domain

type EQuestionMatrix struct {
	OptionsShuffleType    *string `json:"optionsShuffleType"`
	MatrixFormatType      *string `json:"matrixFormatType"`
	MatrixRowNamesAlign   *string `json:"matrixRowNamesAlign"`
	AnswerLimitType       *string `json:"answerLimitType"`
	OptionsShuffle        *bool   `json:"optionsShuffle"`
	OptionsMultipleSelect *bool   `json:"optionsMultipleSelect"`
	ShowAsClosed          *bool   `json:"showAsClosed"`
	AllRowsRequired       *bool   `json:"allRowsRequired"`
	ShowRowsSequentially  *bool   `json:"showRowsSequentially"`
	EnableImages          *bool   `json:"enableImages"`
}

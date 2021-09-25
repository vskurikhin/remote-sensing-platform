package domain

type EQuestionCsi struct {
	RevertOptionsNumeration *bool   `json:"revertOptionsNumeration"`
	ValueMax                *int    `json:"valueMax"`
	ValueSatisfactory       *int    `json:"valueSatisfactory"`
	StartFromOne            *bool   `json:"startFromOne"`
	GradientEnabled         *bool   `json:"gradientEnabled"`
	GradientColorLeft       *string `json:"gradientColorLeft"`
	GradientColorRight      *string `json:"gradientColorRight"`
	Shape                   *string `json:"shape"`
	CommentShowType         *string `json:"commentShowType"`
}

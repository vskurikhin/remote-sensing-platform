package domain

type EQuestionNps struct {
	Neutral                 *int    `json:"neutral"`
	Promoter                *int    `json:"promoter"`
	RevertOptionsNumeration *bool   `json:"revertOptionsNumeration"`
	StartFromOne            *bool   `json:"startFromOne"`
	GradientEnabled         *bool   `json:"gradientEnabled"`
	GradientColorLeft       *string `json:"gradientColorLeft"`
	GradientColorRight      *string `json:"gradientColorRight"`
	Shape                   *string `json:"shape"`
}

package dto

import (
	"github.com/google/uuid"
	"github.com/vskurikhin/remote-sensing-platform/constructor/domain"
)

type PollDesignAndSettings struct {
	Id                          int64      `json:"id"`
	UserId                      *int64     `json:"userId"`
	Template                    bool       `json:"template"`
	BackgroundMediaId           *uuid.UUID `json:"backgroundMediaId"`
	LogoId                      *uuid.UUID `json:"logoId"`
	LogoEnabled                 bool       `json:"logoEnabled"`
	Font                        string     `json:"font"`
	ColorQuestion               string     `json:"colorQuestion"`
	ColorDescription            string     `json:"colorDescription"`
	ColorOptions                string     `json:"colorOptions"`
	ColorBackground             string     `json:"colorBackground"`
	ColorButtons                string     `json:"colorButtons"`
	ColorButtonsText            string     `json:"colorButtonsText"`
	ColorBottomPanel            string     `json:"colorBottomPanel"`
	ColorProgressBar            string     `json:"colorProgressBar"`
	ColorBackgroundTransparency string     `json:"colorBackgroundTransparency"`

	DesignId                  int64    `json:"designId"`
	ProgressBarType           string   `json:"progressBarType"`
	QuestionAlignType         string   `json:"questionAlignType"`
	NotificationFrequency     string   `json:"notificationFrequency"`
	HorizontalScroll          bool     `json:"horizontalScroll"`
	WelcomeScreenFooter       bool     `json:"welcomeScreenFooter"`
	Language                  string   `json:"language"`
	ShuffleEnabled            bool     `json:"shuffleEnabled"`
	SinglePage                bool     `json:"singlePage"`
	ButtonText                string   `json:"buttonText"`
	ButtonNextDisabled        bool     `json:"buttonNextDisabled"`
	ButtonPrevEnabled         bool     `json:"buttonPrevEnabled"`
	ButtonResultEnabled       bool     `json:"buttonResultEnabled"`
	QuestionBackgroundEnabled bool     `json:"questionBackgroundEnabled"`
	QuestionNumbersEnabled    bool     `json:"questionNumbersEnabled"`
	QuestionRequiredHighlight bool     `json:"questionRequiredHighlight"`
	SessionTime               int64    `json:"sessionTime"`
	SessionRestore            bool     `json:"sessionRestore"`
	SessionNumberShow         bool     `json:"sessionNumberShow"`
	RequiredParams            []string `json:"requiredParams"`
}

func NewPollDesignAndSettings(design *domain.EPollDesign, settings *domain.EPollSettings) *PollDesignAndSettings {
	return &PollDesignAndSettings{
		Id:                          design.Id,
		UserId:                      design.UserId,
		Template:                    design.Template,
		BackgroundMediaId:           design.BackgroundMediaId,
		LogoId:                      design.LogoId,
		LogoEnabled:                 design.LogoEnabled,
		Font:                        design.Font,
		ColorQuestion:               design.ColorQuestion,
		ColorDescription:            design.ColorDescription,
		ColorOptions:                design.ColorOptions,
		ColorBackground:             design.ColorBackground,
		ColorButtons:                design.ColorButtons,
		ColorButtonsText:            design.ColorButtonsText,
		ColorBottomPanel:            design.ColorBottomPanel,
		ColorProgressBar:            design.ColorProgressBar,
		ColorBackgroundTransparency: design.ColorBackgroundTransparency,

		DesignId:                  settings.DesignId,
		ProgressBarType:           settings.ProgressBarType,
		QuestionAlignType:         settings.QuestionAlignType,
		NotificationFrequency:     settings.NotificationFrequency,
		HorizontalScroll:          settings.HorizontalScroll,
		WelcomeScreenFooter:       settings.WelcomeScreenFooter,
		Language:                  settings.Language,
		ShuffleEnabled:            settings.ShuffleEnabled,
		SinglePage:                settings.SinglePage,
		ButtonText:                settings.ButtonText,
		ButtonNextDisabled:        settings.ButtonNextDisabled,
		ButtonPrevEnabled:         settings.ButtonPrevEnabled,
		ButtonResultEnabled:       settings.ButtonResultEnabled,
		QuestionBackgroundEnabled: settings.QuestionBackgroundEnabled,
		QuestionNumbersEnabled:    settings.QuestionNumbersEnabled,
		QuestionRequiredHighlight: settings.QuestionRequiredHighlight,
		SessionTime:               settings.SessionTime,
		SessionRestore:            settings.SessionRestore,
		SessionNumberShow:         settings.SessionNumberShow,
		RequiredParams:            settings.RequiredParams,
	}
}

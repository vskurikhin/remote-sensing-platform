package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type EPollSettings struct {
	Id                        int64    `json:"id"`
	DesignId                  int64    `json:"designId"`
	ProgressBarType           string   `json:"progressBarType"`
	QuestionAlignType         string   `json:"QuestionAlignType"`
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

func (e *EPollSettings) String() string {
	return string(e.Marshal())
}

func (e *EPollSettings) Marshal() []byte {

	user, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return user
}

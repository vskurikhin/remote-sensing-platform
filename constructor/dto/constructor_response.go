package dto

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/remote-sensing-platform/constructor/domain"
)

type ConstructorResponse struct {
	Poll           *domain.EPoll            `json:"poll"`
	Settings       *PollDesignAndSettings   `json:"settings"`
	PollChannels   []domain.EPollChannel    `json:"pollChannels"`
	ScreenMains    []domain.FScreenMain     `json:"screen_main"`
	ScreenWelcome  []domain.FScreenWelcome  `json:"screen_welcome"`
	ScreenComplete []domain.FScreenComplete `json:"screen_complete"`
}

func NewConstructorResponse(
	poll *domain.EPoll,
	design *domain.EPollDesign,
	settings *domain.EPollSettings,
	channels []domain.EPollChannel,
	screenMain []domain.FScreenMain,
	screenWelcome []domain.FScreenWelcome,
	screenComplete []domain.FScreenComplete) *ConstructorResponse {

	return &ConstructorResponse{
		Poll:           poll,
		Settings:       NewPollDesignAndSettings(design, settings),
		PollChannels:   channels,
		ScreenMains:    screenMain,
		ScreenWelcome:  screenWelcome,
		ScreenComplete: screenComplete,
	}
}

func (e *ConstructorResponse) String() string {
	return string(e.Marshal())
}

func (e *ConstructorResponse) Marshal() []byte {
	result, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

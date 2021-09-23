package domain

import (
	"context"
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type EPollSettings struct {
	Id                        int64    `json:"id"`
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

func (e *EPollSettings) String() string {
	return string(e.Marshal())
}

func (e *EPollSettings) Marshal() []byte {
	result, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

func (d *ePollSettings) FindById(id int64) (*EPollSettings, error) {
	return d.findById(id)
}

const SELECT_POLL_SETTINGS_BY_ID = `
	SELECT ps.id,
		   ps.button_next_disabled,
		   ps.button_prev_enabled,
		   ps.button_result_enabled,
		   ps.button_text,
		   ps.design_id,
		   ps.horizontal_scroll,
		   ps.language,
		   ps.notification_frequency,
		   ps.progress_bar_type,
		   ps.question_align_type,
		   ps.question_background_enabled,
		   ps.question_numbers_enabled,
		   ps.question_required_highlight,
		   ps.required_params,
		   ps.session_number_show,
		   ps.session_restore,
		   ps.session_time,
		   ps.shuffle_enabled,
		   ps.single_page,
		   ps.welcome_screen_footer
	  FROM poll_settings ps
	 WHERE ps.id = $1`

func (d *ePollSettings) findById(id int64) (*EPollSettings, error) {
	var e EPollSettings
	err := d.poolRo.
		QueryRow(context.Background(), SELECT_POLL_SETTINGS_BY_ID, id).
		Scan(&e.Id,
			&e.ButtonNextDisabled,
			&e.ButtonPrevEnabled,
			&e.ButtonResultEnabled,
			&e.ButtonText,
			&e.DesignId,
			&e.HorizontalScroll,
			&e.Language,
			&e.NotificationFrequency,
			&e.ProgressBarType,
			&e.QuestionAlignType,
			&e.QuestionBackgroundEnabled,
			&e.QuestionNumbersEnabled,
			&e.QuestionRequiredHighlight,
			&e.RequiredParams,
			&e.SessionNumberShow,
			&e.SessionRestore,
			&e.SessionTime,
			&e.ShuffleEnabled,
			&e.SinglePage,
			&e.WelcomeScreenFooter)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

package domain

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/savsgio/go-logger/v2"
)

type EPollDesign struct {
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
}

func (e *EPollDesign) String() string {
	return string(e.Marshal())
}

func (e *EPollDesign) Marshal() []byte {
	result, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

func (d *ePollDesign) FindById(id int64) (*EPollDesign, error) {
	return d.findById(id)
}

const SELECT_POLL_DESIGN_BY_ID = `
	SELECT pd.id,
           pd.background_media_id,
           pd.color_background,
           pd.color_background_transparency,
           pd.color_bottom_panel,
           pd.color_buttons,
           pd.color_buttons_text,
           pd.color_description ,
           pd.color_options,
           pd.color_progress_bar,
           pd.color_question,
           pd.font,
           pd.logo_enabled,
           pd.logo_id,
           pd.user_id,
           (SELECT EXISTS(SELECT 1 FROM poll_design_template t WHERE t.id = pd.id)) AS template
      FROM poll_design pd
     WHERE pd.id = $1`

func (d *ePollDesign) findById(id int64) (*EPollDesign, error) {
	var e EPollDesign
	err := d.poolRo.
		QueryRow(context.Background(), SELECT_POLL_DESIGN_BY_ID, id).
		Scan(&e.Id,
			&e.BackgroundMediaId,
			&e.ColorBackground,
			&e.ColorBackgroundTransparency,
			&e.ColorBottomPanel,
			&e.ColorButtons,
			&e.ColorButtonsText,
			&e.ColorDescription,
			&e.ColorOptions,
			&e.ColorProgressBar,
			&e.ColorQuestion,
			&e.Font,
			&e.LogoEnabled,
			&e.LogoId,
			&e.UserId,
			&e.Template)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

package domain

import (
	"context"
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type FScreenWelcome struct {
	PollItem      *EPollItem      `json:"pollItem"`
	ScreenWelcome *EScreenWelcome `json:"screenWelcome"`
}

func (e *FScreenWelcome) String() string {
	return string(e.Marshal())
}

func (e *FScreenWelcome) Marshal() []byte {
	result, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

func (d *fScreenWelcome) FindById(id int64) ([]FScreenWelcome, error) {
	return d.findById(id)
}

const SELECT_SCREEN_WELCOME_BY_ID = `
	SELECT pi.id,
           pi.custom_shape_set_id,
           pi.default_screen_out,
           pi.default_transition_id,
           pi.image_size_type,
           pi.media_location_type,
           pi.poll_id,
           pi.show_logic_id,
           pi.show_type,
           pi.type,
           sw.button_text,
           sw.description,
           sw.index,
           sw.name
      FROM poll_item pi
      LEFT OUTER JOIN screen_welcome sw ON pi.id = sw.id
     WHERE pi.type = 'WELCOME_SCREEN'
       AND pi.poll_id = $1
       AND (pi.deleted_at IS NULL)
     ORDER by sw.index ASC`

func (d *fScreenWelcome) findById(id int64) ([]FScreenWelcome, error) {
	var list []FScreenWelcome
	rows, err := d.poolRo.Query(context.Background(), SELECT_SCREEN_WELCOME_BY_ID, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		e := FScreenWelcome{
			PollItem:      new(EPollItem),
			ScreenWelcome: new(EScreenWelcome),
		}
		err := rows.Scan(&e.PollItem.Id,
			&e.PollItem.CustomShapeSetId,
			&e.PollItem.DefaultScreenOut,
			&e.PollItem.DefaultTransitionId,
			&e.PollItem.ImageSizeType,
			&e.PollItem.MediaLocationType,
			&e.PollItem.PollId,
			&e.PollItem.ShowLogicId,
			&e.PollItem.ShowType,
			&e.PollItem.Type,
			&e.ScreenWelcome.ButtonText,
			&e.ScreenWelcome.Description,
			&e.ScreenWelcome.Index,
			&e.ScreenWelcome.Name)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}

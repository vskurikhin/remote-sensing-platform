package domain

import (
	"context"
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type FScreenComplete struct {
	PollItem       *EPollItem       `json:"pollItem"`
	ScreenComplete *EScreenComplete `json:"screenComplete"`
}

func (e *FScreenComplete) String() string {
	return string(e.Marshal())
}

func (e *FScreenComplete) Marshal() []byte {
	result, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

func (d *fScreenComplete) FindById(id int64) ([]FScreenComplete, error) {
	return d.findById(id)
}

const SELECT_SCREEN_COMPLETE_BY_ID = `
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
       sc.button_text,
       sc.description,
       sc.index,
       sc.link,
       sc.link_facebook,
       sc.link_in_same_screen,
       sc.link_instagram,
       sc.link_telegram,
       sc.link_twitter,
       sc.link_vk,
       sc.my_answer_button_text,
       sc.my_answer_enabled,
       sc.name
  FROM poll_item pi
  LEFT OUTER JOIN screen_complete sc ON pi.id = sc.id
 WHERE pi.type = 'COMPLETE_SCREEN'
   AND pi.poll_id = $1
   AND (pi.deleted_at IS NULL)
 ORDER by sc.index ASC`

func (d *fScreenComplete) findById(id int64) ([]FScreenComplete, error) {
	var list []FScreenComplete
	rows, err := d.poolRo.Query(context.Background(), SELECT_SCREEN_COMPLETE_BY_ID, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		e := FScreenComplete{
			PollItem:       new(EPollItem),
			ScreenComplete: new(EScreenComplete),
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
			&e.ScreenComplete.ButtonText,
			&e.ScreenComplete.Description,
			&e.ScreenComplete.Index,
			&e.ScreenComplete.Link,
			&e.ScreenComplete.LinkFacebook,
			&e.ScreenComplete.LinkInSameScreen,
			&e.ScreenComplete.LinkInstagram,
			&e.ScreenComplete.LinkTelegram,
			&e.ScreenComplete.LinkTwitter,
			&e.ScreenComplete.LinkVk,
			&e.ScreenComplete.MyAnswerButtonText,
			&e.ScreenComplete.MyAnswerEnabled,
			&e.ScreenComplete.Name)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}

package domain

import (
	"context"
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
	"time"
)

type EPollChannel struct {
	Id                 int64      `json:"id"`
	PollId             int64      `json:"pollId"`
	OneAnswerPerDevice bool       `json:"oneAnswerPerDevice"`
	HashUniqueAnswer   bool       `json:"hashUniqueAnswer"`
	HashAnswerMax      int        `json:"hashAnswerMax"`
	Enabled            bool       `json:"enabled"`
	ShowShortCode      bool       `json:"showShortCode"`
	Code               string     `json:"code"`
	ShortCode          string     `json:"shortCode"`
	Type               string     `json:"type"`
	Name               string     `json:"name"`
	MaxAnswers         *int       `json:"maxAnswers"`
	StopTimestamp      *time.Time `json:"stopTimestamp"`
	CompleteLink       *string    `json:"completeLink"`
	ScreenoutLink      *string    `json:"screenoutLink"`
	QuotafullLink      *string    `json:"quotafullLink"`
	Quotas             []string   `json:"quotas"` // TODO refactoring
	CustomId           *string    `json:"customId"`
	Debug              bool       `json:"debug"`
}

func (e *EPollChannel) String() string {
	return string(e.Marshal())
}

func (e *EPollChannel) Marshal() []byte {
	result, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

func (d *ePollChannel) FindAllByPollIdAndDebugFalseOrderByIdAsc(id int64) ([]EPollChannel, error) {
	return d.findAllByPollIdAndDebugFalseOrderByIdAsc(id)
}

const SELECT_POLL_CHANNEL_BY_ID = `
	SELECT pc.id,
		   pc.code,
		   pc.complete_link,
		   pc.custom_id,
		   pc.enabled,
		   pc.hash_answer_max,
		   pc.hash_unique_answer,
		   pc.max_answers,
		   pc.name,
		   pc.one_answer_per_device,
		   pc.poll_id,
		   pc.quotafull_link,
		   pc.screenout_link,
		   pc.short_code,
		   pc.show_short_code,
		   pc.stop_timestamp,
		   pc.type,
		   (pc.type = 'DEBUG') AS debug
	  FROM poll_channel pc
	 WHERE (pc.deleted_at IS NULL)
	   AND pc.poll_id = $1
	   AND (pc.type = 'DEBUG') = false
	 ORDER BY pc.id ASC`

func (d *ePollChannel) findAllByPollIdAndDebugFalseOrderByIdAsc(id int64) ([]EPollChannel, error) {
	var list []EPollChannel
	rows, err := d.poolRo.Query(context.Background(), SELECT_POLL_CHANNEL_BY_ID, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var e EPollChannel
		err := rows.Scan(&e.Id,
			&e.Code,
			&e.CompleteLink,
			&e.CustomId,
			&e.Enabled,
			&e.HashAnswerMax,
			&e.HashUniqueAnswer,
			&e.MaxAnswers,
			&e.Name,
			&e.OneAnswerPerDevice,
			&e.PollId,
			&e.QuotafullLink,
			&e.ScreenoutLink,
			&e.ShortCode,
			&e.ShowShortCode,
			&e.StopTimestamp,
			&e.Type,
			&e.Debug)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}

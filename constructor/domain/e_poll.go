package domain

import (
	"context"
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
	"time"
)

type EPoll struct {
	Id               int64     `json:"id"`
	SegmentId        int64     `json:"segmentId"`
	UserId           int64     `json:"userId"`
	FolderId         *int64    `json:"folderId"`
	ExportSettingsId *int64    `json:"exportSettingsId"`
	Name             string    `json:"name"`
	Status           string    `json:"status"`
	Type             string    `json:"type"`
	Tags             []string  `json:"tags"`
	Favorite         bool      `json:"favorite"`
	InTrash          bool      `json:"inTrash"`
	CreatedAt        time.Time `json:"createdAt"`
	AnswersCount     int64     `json:"answersCount"`
	Owner            EUser     `json:"owner"`
}

func (e *EPoll) String() string {
	return string(e.Marshal())
}

func (e *EPoll) Marshal() []byte {
	result, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

func (d *ePoll) FindById(id int64) (*EPoll, error) {
	return d.findById(id)
}

const SELECT_POLL_BY_ID = `
	SELECT ep.id,
		   ep.created_at,
		   ep.export_settings_id,
		   ep.favorite,
		   ep.folder_id,
		   ep.in_trash,
		   ep.name,
		   ep.user_id,
		   ep.segment_id,
		   ep.status,
		   ep.tags,
		   ep.type,
		   (SELECT count(a.id)
			  FROM answer_poll a
			 WHERE a.segment_id = ep.segment_id
			   AND a.poll_id = ep.id
			   AND a.deleted_at IS NULL
			   AND a.finished_at IS NOT NULL
			   AND a.debug = false
			   AND a.status = 'FINISHED'
			   AND a.processing_status != 'EXCLUDED') as answersCount
      FROM poll ep
     WHERE ep.id = $1
       AND (ep.deleted_at IS NULL)`

func (d *ePoll) findById(id int64) (*EPoll, error) {
	var e EPoll
	err := d.poolRo.
		QueryRow(context.Background(), SELECT_POLL_BY_ID, id).
		Scan(&e.Id,
			&e.CreatedAt,
			&e.ExportSettingsId,
			&e.Favorite,
			&e.FolderId,
			&e.InTrash,
			&e.Name,
			&e.UserId,
			&e.SegmentId,
			&e.Status,
			&e.Tags,
			&e.Type,
			&e.AnswersCount)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

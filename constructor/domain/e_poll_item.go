package domain

type EPollItem struct {
	Id                  int64    `json:"id"`
	PollId              int64    `json:"pollId"`
	Type                *string  `json:"type"`
	ShowType            *string  `json:"showType"`
	DefaultTransitionId *int64   `json:"defaultTransitionId"`
	DefaultScreenOut    *bool    `json:"defaultScreenOut"`
	ImageSizeType       *string  `json:"imageSizeType"`
	MediaLocationType   *string  `json:"mediaLocationType"`
	CustomShapeSetId    *int64   `json:"customShapeSetId"`
	ShowLogic           *string  `json:"showLogic"` // TODO
	Media               []string `json:"media"`
	Transitions         []string `json:"transitions"`
	CustomShapeSet      *string  `json:"customShapeSet"`
	ShowLogicId         *int64   `json:"showLogicId"`
}

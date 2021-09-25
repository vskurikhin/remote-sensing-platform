package domain

type EQuestion struct {
	Name                map[string]interface{} `json:"name"`
	Description         map[string]interface{} `json:"description"`
	ButtonText          map[string]interface{} `json:"buttonText"`
	NotificationText    map[string]interface{} `json:"notificationText"`
	CommentPlaceholder  map[string]interface{} `json:"commentPlaceholder"`
	DuplicateId         *int64                 `json:"duplicateId"`
	DisplayIndex        *string                `json:"displayIndex"`
	Required            *bool                  `json:"required"`
	NotificationEnabled *bool                  `json:"notificationEnabled"`
	CommentEnabled      *bool                  `json:"commentEnabled"`
	HasAnswers          *bool                  `json:"hasAnswers"`
	SystemName          *string                `json:"systemName"`
	Options             []string               `json:"options"` // TODO
	Parents             []string               `json:"parents"` // TODO
}

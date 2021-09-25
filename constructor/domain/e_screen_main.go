package domain

type EScreenMain struct {
	ParentId   *int64 `json:"parentId"`
	Index      int64  `json:"index"`
	Pin        bool   `json:"pin"`
	LocalIndex int64  `json:"localIndex"`
}

package domain

type EQuestionComparison struct {
	MediaEnabled       *bool `json:"mediaEnabled"`
	OptionsShuffle     *bool `json:"optionsShuffle"`
	ChangeAnswer       *bool `json:"changeAnswer"`
	ShowRemainingPairs *bool `json:"showRemainingPairs"`
}

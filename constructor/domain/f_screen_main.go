package domain

import (
	"context"
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
)

type FScreenMain struct {
	PollItem                     *EPollItem                     `json:"pollItem"`
	ScreenMain                   *EScreenMain                   `json:"screenMain"`
	Group                        *EGroup                        `json:"group"`
	Question                     *EQuestion                     `json:"question"`
	QuestionAudio                *EQuestionAudio                `json:"questionAudio"`
	QuestionCardSorting          *EQuestionCardSorting          `json:"questionCardSorting"`
	QuestionClosed               *EQuestionClosed               `json:"questionClosed"`
	QuestionComparison           *EQuestionComparison           `json:"questionComparison"`
	QuestionCsi                  *EQuestionCsi                  `json:"questionCsi"`
	QuestionFirstClick           *EQuestionFirstClick           `json:"questionFirstClick"`
	QuestionMatrix               *EQuestionMatrix               `json:"questionMatrix"`
	QuestionMedia                *EQuestionMedia                `json:"questionMedia"`
	QuestionNps                  *EQuestionNps                  `json:"questionNps"`
	QuestionOpened               *EQuestionOpened               `json:"questionOpened"`
	QuestionOpinion              *EQuestionOpinion              `json:"questionOpinion"`
	QuestionPassword             *EQuestionPassword             `json:"questionPassword"`
	QuestionRanging              *EQuestionRanging              `json:"questionRanging"`
	QuestionRating               *EQuestionRating               `json:"questionRating"`
	QuestionSemanticDifferential *EQuestionSemanticDifferential `json:"questionSemanticDifferential"`
	QuestionSiteTest             *EQuestionSiteTest             `json:"questionSiteTest"`
	QuestionSlideshow            *EQuestionSlideshow            `json:"questionSlideshow"`
	QuestionTreeTesting          *EQuestionTreeTesting          `json:"questionTreeTesting"`
}

func (e *FScreenMain) String() string {
	return string(e.Marshal())
}

func (e *FScreenMain) Marshal() []byte {
	result, err := json.Marshal(*e)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return result
}

func (d *fPollItem) FindById(id int64) ([]FScreenMain, error) {
	return d.findById(id)
}

const SELECT_POLL_ITEM_BY_ID = `
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
       sm.index,
       sm.parent_id,
       sm.pin,
       gr.button_text,
       gr.indent,
       gr.name,
       gr.questions_shuffle,
       q00.button_text,
       q00.comment_enabled,
       q00.comment_placeholder,
       q00.description,
       q00.display_index,
       q00.duplicate_id,
       q00.name,
       q00.notification_enabled,
       q00.notification_text,
       q00.required,
       q00.system_name,
       qau.min_listening_time,
       qc1.shuffle_cards,
       qc1.shuffle_category,
       qcl.options_horizontal,
       qcl.options_index_type,
       qcl.options_multiple_select,
       qcl.options_multiple_select_max,
       qcl.options_multiple_select_min,
       qcl.options_show_type,
       qcl.options_shuffle,
       qcl.options_sort_alphabetically,
       qcm.change_answer,
       qcm.media_enabled,
       qcm.options_shuffle,
       qcm.show_remaining_pairs,
       qc2.comment_show_type,
       qc2.gradient_color_left,
       qc2.gradient_color_right,
       qc2.gradient_enabled,
       qc2.revert_options_numeration,
       qc2.shape,
       qc2.start_from_one,
       qc2.value_max,
       qc2.value_satisfactory,
       qfc.hash,
       qmx.all_rows_required,
       qmx.answer_limit_type,
       qmx.enable_images,
       qmx.matrix_format_type,
       qmx.matrix_row_names_align,
       qmx.options_multiple_select,
       qmx.options_shuffle,
       qmx.options_shuffle_type,
       qmx.show_as_closed,
       qmx.show_rows_sequentially,
       qmd.options_full_screen_enabled,
       qmd.options_increased,
       qmd.options_landscape_format,
       qmd.options_multiple_select,
       qmd.options_multiple_select_max,
       qmd.options_multiple_select_min,
       qmd.options_select_button_disabled,
       qmd.options_shuffle,
       qnp.gradient_color_left,
       qnp.gradient_color_right,
       qnp.gradient_enabled,
       qnp.neutral,
       qnp.promoter,
       qnp.revert_options_numeration,
       qnp.shape,
       qnp.start_from_one,
       qo1.hide_tip,
       qo1.multiline,
       qo1.options_multiple_select_min,
       qo1.sequential_display,
       qo1.string_max_length,
       qo1.string_min_length,
       qo2.emotion_scale,
       qo2.gradient_color_left,
       qo2.gradient_color_right,
       qo2.gradient_enabled,
       qo2.options_value_max,
       qo2.options_value_min,
       qo2.shape,
       qpw.login_identifier_pairs,
       qpw.login_placeholder,
       qpw.password,
       qpw.password_placeholder,
       qr1.options_shuffle,
       qr1.show_options_numeration,
       qr2.color,
       qr2.gradient_color_left,
       qr2.gradient_color_right,
       qr2.gradient_enabled,
       qr2.option_initial_interest_value,
       qr2.option_initial_number_value,
       qr2.options_value_max,
       qr2.options_value_min,
       qr2.revert_options_numeration,
       qr2.scale,
       qr2.section_count,
       qr2.shape,
       qr2.show_options_numeration,
       qr2.slider_view_type,
       qr2.start_from_one,
       qr2.step,
       qr2.sub_type,
       qsd.gradient_color_left,
       qsd.gradient_color_right,
       qsd.gradient_enabled,
       qsd.hide_zero,
       qsd.options_shuffle,
       qsd.shape,
       qsd.show_options_numeration,
       qst.hash,
       qst.version,
       qss.show_time,
       qtr.count_steps,
       qtr.time
  FROM poll_item pi
  LEFT OUTER JOIN screen_main sm                     ON pi.id = sm.id
  LEFT OUTER JOIN public.group gr                    ON pi.id = gr.id
  LEFT OUTER JOIN question q00                       ON pi.id = q00.id
  LEFT OUTER JOIN question_audio qau                 ON pi.id = qau.id
  LEFT OUTER JOIN question_card_sorting qc1          ON pi.id = qc1.id
  LEFT OUTER JOIN question_closed qcl                ON pi.id = qcl.id
  LEFT OUTER JOIN question_comparison qcm            ON pi.id = qcm.id
  LEFT OUTER JOIN question_csi qc2                   ON pi.id = qc2.id
  LEFT OUTER JOIN question_first_click qfc           ON pi.id = qfc.id
  LEFT OUTER JOIN question_matrix qmx                ON pi.id = qmx.id
  LEFT OUTER JOIN question_media qmd                 ON pi.id = qmd.id
  LEFT OUTER JOIN question_nps qnp                   ON pi.id = qnp.id
  LEFT OUTER JOIN question_opened qo1                ON pi.id = qo1.id
  LEFT OUTER JOIN question_opinion qo2               ON pi.id = qo2.id
  LEFT OUTER JOIN question_password qpw              ON pi.id = qpw.id
  LEFT OUTER JOIN question_ranging qr1               ON pi.id = qr1.id
  LEFT OUTER JOIN question_rating qr2                ON pi.id = qr2.id
  LEFT OUTER JOIN question_semantic_differential qsd ON pi.id = qsd.id
  LEFT OUTER JOIN question_site_test qst             ON pi.id = qst.id
  LEFT OUTER JOIN question_slideshow qss             ON pi.id = qss.id
  LEFT OUTER JOIN question_test_text qtt             ON pi.id = qtt.id
  LEFT OUTER JOIN question_tree_testing qtr          ON pi.id = qtr.id
 WHERE pi.type IN
      ('AGREEMENT', 'AUDIO', 'CARD_SORTING', 'CLOSED', 'COMPARISON', 'CSI', 'FILE', 'FIRST_CLICK', 'KANO',
       'LINKED_LIST', 'MATRIX', 'MEDIA', 'NPS', 'OPENED', 'OPINION', 'PASSWORD', 'RANGING', 'RATING',
       'SEMANTIC_DIFFERENTIAL', 'SITE_TEST', 'SLIDESHOW', 'TEST_TEXT', 'TEXT', 'TREE_TESTING', 'GROUP')
   AND (pi.deleted_at IS NULL)
   AND pi.poll_id = $1
   AND (pi.deleted_at IS NULL)
 ORDER BY sm.index ASC
`

func (d *fPollItem) findById(id int64) ([]FScreenMain, error) {
	var list []FScreenMain
	rows, err := d.poolRo.Query(context.Background(), SELECT_POLL_ITEM_BY_ID, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		e := FScreenMain{
			PollItem:                     new(EPollItem),
			ScreenMain:                   new(EScreenMain),
			Group:                        new(EGroup),
			Question:                     new(EQuestion),
			QuestionAudio:                new(EQuestionAudio),
			QuestionCardSorting:          new(EQuestionCardSorting),
			QuestionClosed:               new(EQuestionClosed),
			QuestionComparison:           new(EQuestionComparison),
			QuestionCsi:                  new(EQuestionCsi),
			QuestionFirstClick:           new(EQuestionFirstClick),
			QuestionMatrix:               new(EQuestionMatrix),
			QuestionMedia:                new(EQuestionMedia),
			QuestionNps:                  new(EQuestionNps),
			QuestionOpened:               new(EQuestionOpened),
			QuestionOpinion:              new(EQuestionOpinion),
			QuestionPassword:             new(EQuestionPassword),
			QuestionRanging:              new(EQuestionRanging),
			QuestionRating:               new(EQuestionRating),
			QuestionSemanticDifferential: new(EQuestionSemanticDifferential),
			QuestionSiteTest:             new(EQuestionSiteTest),
			QuestionSlideshow:            new(EQuestionSlideshow),
			QuestionTreeTesting:          new(EQuestionTreeTesting),
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
			&e.ScreenMain.Index,
			&e.ScreenMain.ParentId,
			&e.ScreenMain.Pin,
			&e.Group.ButtonText,
			&e.Group.Indent,
			&e.Group.Name,
			&e.Group.QuestionsShuffle,
			&e.Question.ButtonText,
			&e.Question.CommentEnabled,
			&e.Question.CommentPlaceholder,
			&e.Question.Description,
			&e.Question.DisplayIndex,
			&e.Question.DuplicateId,
			&e.Question.Name,
			&e.Question.NotificationEnabled,
			&e.Question.NotificationText,
			&e.Question.Required,
			&e.Question.SystemName,
			&e.QuestionAudio.MinListeningTime,
			&e.QuestionCardSorting.ShuffleCards,
			&e.QuestionCardSorting.ShuffleCategory,
			&e.QuestionClosed.OptionsHorizontal,
			&e.QuestionClosed.OptionsIndexType,
			&e.QuestionClosed.OptionsMultipleSelect,
			&e.QuestionClosed.OptionsMultipleSelectMax,
			&e.QuestionClosed.OptionsMultipleSelectMin,
			&e.QuestionClosed.OptionsShowType,
			&e.QuestionClosed.OptionsShuffle,
			&e.QuestionClosed.OptionsSortAlphabetically,
			&e.QuestionComparison.ChangeAnswer,
			&e.QuestionComparison.MediaEnabled,
			&e.QuestionComparison.OptionsShuffle,
			&e.QuestionComparison.ShowRemainingPairs,
			&e.QuestionCsi.CommentShowType,
			&e.QuestionCsi.GradientColorRight,
			&e.QuestionCsi.GradientColorRight,
			&e.QuestionCsi.GradientEnabled,
			&e.QuestionCsi.RevertOptionsNumeration,
			&e.QuestionCsi.Shape,
			&e.QuestionCsi.StartFromOne,
			&e.QuestionCsi.ValueMax,
			&e.QuestionCsi.ValueSatisfactory,
			&e.QuestionFirstClick.Hash,
			&e.QuestionMatrix.AllRowsRequired,
			&e.QuestionMatrix.AnswerLimitType,
			&e.QuestionMatrix.EnableImages,
			&e.QuestionMatrix.MatrixFormatType,
			&e.QuestionMatrix.MatrixRowNamesAlign,
			&e.QuestionMatrix.OptionsMultipleSelect,
			&e.QuestionMatrix.OptionsShuffle,
			&e.QuestionMatrix.OptionsShuffleType,
			&e.QuestionMatrix.ShowAsClosed,
			&e.QuestionMatrix.ShowRowsSequentially,
			&e.QuestionMedia.OptionsFullScreenEnabled,
			&e.QuestionMedia.OptionsIncreased,
			&e.QuestionMedia.OptionsLandscapeFormat,
			&e.QuestionMedia.OptionsMultipleSelect,
			&e.QuestionMedia.OptionsMultipleSelectMax,
			&e.QuestionMedia.OptionsMultipleSelectMin,
			&e.QuestionMedia.OptionsSelectButtonDisabled,
			&e.QuestionMedia.OptionsShuffle,
			&e.QuestionNps.GradientColorLeft,
			&e.QuestionNps.GradientColorRight,
			&e.QuestionNps.GradientEnabled,
			&e.QuestionNps.Neutral,
			&e.QuestionNps.Promoter,
			&e.QuestionNps.RevertOptionsNumeration,
			&e.QuestionNps.Shape,
			&e.QuestionNps.StartFromOne,
			&e.QuestionOpened.HideTip,
			&e.QuestionOpened.Multiline,
			&e.QuestionOpened.OptionsMultipleSelectMin,
			&e.QuestionOpened.SequentialDisplay,
			&e.QuestionOpened.StringMaxLength,
			&e.QuestionOpened.StringMinLength,
			&e.QuestionOpinion.EmotionScale,
			&e.QuestionOpinion.GradientColorLeft,
			&e.QuestionOpinion.GradientColorRight,
			&e.QuestionOpinion.GradientEnabled,
			&e.QuestionOpinion.OptionsValueMax,
			&e.QuestionOpinion.OptionsValueMin,
			&e.QuestionOpinion.Shape,
			&e.QuestionPassword.LoginIdentifierPairs,
			&e.QuestionPassword.LoginPlaceholder,
			&e.QuestionPassword.Password,
			&e.QuestionPassword.PasswordPlaceholder,
			&e.QuestionRanging.OptionsShuffle,
			&e.QuestionRanging.ShowOptionsNumeration,
			&e.QuestionRating.Color,
			&e.QuestionRating.GradientColorLeft,
			&e.QuestionRating.GradientColorRight,
			&e.QuestionRating.GradientEnabled,
			&e.QuestionRating.OptionInitialInterestValue,
			&e.QuestionRating.OptionInitialNumberValue,
			&e.QuestionRating.OptionsValueMax,
			&e.QuestionRating.OptionsValueMin,
			&e.QuestionRating.RevertOptionsNumeration,
			&e.QuestionRating.Scale,
			&e.QuestionRating.SectionCount,
			&e.QuestionRating.Shape,
			&e.QuestionRating.ShowOptionsNumeration,
			&e.QuestionRating.SliderViewType,
			&e.QuestionRating.StartFromOne,
			&e.QuestionRating.Step,
			&e.QuestionRating.SubType,
			&e.QuestionSemanticDifferential.GradientColorLeft,
			&e.QuestionSemanticDifferential.GradientColorRight,
			&e.QuestionSemanticDifferential.GradientEnabled,
			&e.QuestionSemanticDifferential.HideZero,
			&e.QuestionSemanticDifferential.OptionsShuffle,
			&e.QuestionSemanticDifferential.Shape,
			&e.QuestionSemanticDifferential.ShowOptionsNumeration,
			&e.QuestionSiteTest.Hash,
			&e.QuestionSiteTest.Version,
			&e.QuestionSlideshow.ShowTime,
			&e.QuestionTreeTesting.CountSteps,
			&e.QuestionTreeTesting.Time)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}

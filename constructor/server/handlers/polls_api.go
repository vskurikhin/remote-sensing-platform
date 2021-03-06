package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/remote-sensing-platform/constructor/domain"
	"github.com/vskurikhin/remote-sensing-platform/constructor/dto"
	"strconv"
)

func (h *Handlers) GetPollConstructorPageData(ctx *sa.RequestCtx) error {
	ePollSettings, err := h.getPollConstructorPageData(ctx)
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(ePollSettings.String())
}

func (h *Handlers) getPollConstructorPageData(ctx *sa.RequestCtx) (*dto.ConstructorResponse, error) {

	pollId := fmt.Sprintf("%v", ctx.UserValue("pollId"))
	id, err := strconv.ParseInt(pollId, 10, 64)
	if err != nil {
		return nil, err
	}

	poll, err := h.getEPoll(id)
	if err != nil {
		return nil, err
	}

	settings, err := h.getEPollSettings(id)
	if err != nil {
		return nil, err
	}

	design, err := h.getEPollDesign(settings.DesignId)
	if err != nil {
		return nil, err
	}

	channels, err := h.getAllEPollChannelsByPollIdAndDebugFalseOrderByIdAsc(id)
	if err != nil {
		return nil, err
	}

	screenMain, err := h.getAllFScreenMainByPollIdAndDeletedAtIsNullOrderByIndexWithLocalIndex(id)
	if err != nil {
		return nil, err
	}

	screenWelcome, err := h.getAllFScreenWelcomeByPollIdAndDeletedAtIsNullOrderByIndexWithLocalIndex(id)
	if err != nil {
		return nil, err
	}

	screenComplete, err := h.getAllFScreenCompleteByPollIdAndDeletedAtIsNullOrderByIndexWithLocalIndex(id)
	if err != nil {
		return nil, err
	}

	constructorResponse := dto.NewConstructorResponse(poll, design, settings, channels, screenMain, screenWelcome, screenComplete)

	return constructorResponse, nil
}

func (h *Handlers) getEPoll(id int64) (*domain.EPoll, error) {
	poll, err := h.Server.Cache.GetEPoll(id)
	if err != nil {
		poll, err = h.Server.Dao.EPoll.FindById(id)
		if err != nil {
			return nil, err
		}
		h.Server.Cache.PutEPoll(id, poll)
	}
	return poll, nil
}

func (h *Handlers) getEPollSettings(id int64) (*domain.EPollSettings, error) {
	settings, err := h.Server.Cache.GetEPollSettings(id)
	if err != nil {
		settings, err = h.Server.Dao.EPollSettings.FindById(id)
		if err != nil {
			return nil, err
		}
		h.Server.Cache.PutEPollSettings(id, settings)
	}
	return settings, nil
}

func (h *Handlers) getEPollDesign(id int64) (*domain.EPollDesign, error) {
	design, err := h.Server.Cache.GetEPollDesign(id)
	if err != nil {
		design, err = h.Server.Dao.EPollDesign.FindById(id)
		if err != nil {
			return nil, err
		}
		h.Server.Cache.PutEPollDesign(id, design)
	}
	return design, nil
}

func (h *Handlers) getAllEPollChannelsByPollIdAndDebugFalseOrderByIdAsc(id int64) ([]domain.EPollChannel, error) {
	channels, err := h.Server.Cache.GetEPollChannel(id)
	if err != nil {
		channels, err = h.Server.Dao.EPollChannel.FindAllByPollIdAndDebugFalseOrderByIdAsc(id)
		if err != nil {
			return nil, err
		}
		h.Server.Cache.PutEPollChannel(id, channels)
	}
	return channels, nil
}

func (h *Handlers) getAllFScreenMainByPollIdAndDeletedAtIsNullOrderByIndexWithLocalIndex(id int64) ([]domain.FScreenMain, error) {
	screenMain, err := h.Server.Cache.GetArrayOfFScreenMain(id)
	if screenMain == nil || err != nil {
		screenMain, err = h.Server.Dao.FScreenMain.FindById(id)
		if err != nil {
			return nil, err
		}
		h.Server.Cache.PutArrayOfFScreenMain(id, screenMain)
	}
	return screenMain, nil
}

func (h *Handlers) getAllFScreenWelcomeByPollIdAndDeletedAtIsNullOrderByIndexWithLocalIndex(id int64) ([]domain.FScreenWelcome, error) {
	screenWelcome, err := h.Server.Cache.GetArrayOfFScreenWelcome(id)
	if screenWelcome == nil || err != nil {
		screenWelcome, err = h.Server.Dao.FScreenWelcome.FindById(id)
		if err != nil {
			return nil, err
		}
		h.Server.Cache.PutArrayOfFScreenWelcome(id, screenWelcome)
	}
	return screenWelcome, nil
}

func (h *Handlers) getAllFScreenCompleteByPollIdAndDeletedAtIsNullOrderByIndexWithLocalIndex(id int64) ([]domain.FScreenComplete, error) {
	screenComplete, err := h.Server.Cache.GetArrayOfFScreenComplete(id)
	if screenComplete == nil || err != nil {
		screenComplete, err = h.Server.Dao.FScreenComplete.FindById(id)
		if err != nil {
			return nil, err
		}
		h.Server.Cache.PutArrayOfFScreenComplete(id, screenComplete)
	}
	return screenComplete, nil
}

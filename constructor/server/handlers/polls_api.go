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
	constructorResponse := dto.NewConstructorResponse(poll, design, settings, channels)
	return constructorResponse, nil
}

func (h *Handlers) getEPoll(id int64) (*domain.EPoll, error) {
	return h.Server.Dao.EPoll.FindById(id)
}

func (h *Handlers) getEPollSettings(id int64) (*domain.EPollSettings, error) {
	return h.Server.Dao.EPollSettings.FindById(id)
}

func (h *Handlers) getEPollDesign(id int64) (*domain.EPollDesign, error) {
	return h.Server.Dao.EPollDesign.FindById(id)
}

func (h *Handlers) getAllEPollChannelsByPollIdAndDebugFalseOrderByIdAsc(id int64) ([]domain.EPollChannel, error) {
	return h.Server.Dao.EPollChannel.FindAllByPollIdAndDebugFalseOrderByIdAsc(id)
}

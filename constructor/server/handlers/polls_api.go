package handlers

import (
	"fmt"
	sa "github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
	"github.com/vskurikhin/remote-sensing-platform/constructor/domain"
	"strconv"
)

func (h *Handlers) GetPollConstructorPageData(ctx *sa.RequestCtx) error {
	result, err := h.getPollConstructorPageData(ctx)
	if err != nil {
		logger.Error(err)
		errorCase := domain.ApiMessage{
			Code:    fasthttp.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return ctx.HTTPResponse(errorCase.String(), fasthttp.StatusPreconditionFailed)
	}
	return ctx.HTTPResponse(result.String())
}

func (h *Handlers) getPollConstructorPageData(ctx *sa.RequestCtx) (*domain.EPollSettings, error) {
	pollId := fmt.Sprintf("%v", ctx.UserValue("pollId"))
	id, err := strconv.ParseInt(pollId, 10, 64)
	if err != nil {
		return nil, err
	}
	return h.Server.Dao.EPollSettings.FindById(id)
}

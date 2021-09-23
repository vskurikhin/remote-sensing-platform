package handlers

import (
	sa "github.com/savsgio/atreugo/v11"
	"github.com/vskurikhin/remote-sensing-platform/constructor/domain"
)

func (h *Handlers) Profile(ctx *sa.RequestCtx) error {
	return ctx.HTTPResponse(domain.Dummy.String(), domain.Dummy.Code)
}

package domain

import (
	"encoding/json"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
)

type ApiMessage struct {
	Code    int
	Message string
}

var Dummy = ApiMessage{
	Code:    fasthttp.StatusNotImplemented,
	Message: "dummy",
}

func (a *ApiMessage) String() string {
	return string(a.Marshal())
}

func (a *ApiMessage) Marshal() []byte {
	apiMessage, err := json.Marshal(*a)
	if err != nil {
		logger.Errorf("%v", err)
		return nil
	}
	return apiMessage
}

package endpoint

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/context/endpoint"
	"github.com/mmuflih/rest-api-doc/context/user"
)

/*
 * Getd by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-28 06:18
 */

type GetHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type getHandler struct {
	uc   endpoint.GetUsecase
	rr   httplib.RequestReader
	auth user.GetAuthUserUsecase
}

func (ch getHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := ch.rr.GetRouteParam(r, "document_id")
	err, resp := ch.uc.Get(id)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewGetHandler(uc endpoint.GetUsecase, rr httplib.RequestReader,
	auth user.GetAuthUserUsecase) GetHandler {
	return &getHandler{uc, rr, auth}
}

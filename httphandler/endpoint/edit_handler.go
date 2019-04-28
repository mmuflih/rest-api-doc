package endpoint

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/context/endpoint"
	"github.com/mmuflih/rest-api-doc/context/user"
)

/*
 * Editd by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-28 06:18
 */

type EditHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type editHandler struct {
	uc   endpoint.EditUsecase
	rr   httplib.RequestReader
	auth user.GetAuthUserUsecase
}

func (ch editHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := createRequest{}
	err := ch.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := ch.uc.Edit(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewEditHandler(uc endpoint.EditUsecase, rr httplib.RequestReader,
	auth user.GetAuthUserUsecase) EditHandler {
	return &editHandler{uc, rr, auth}
}

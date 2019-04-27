package user

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/context/user"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 08:52
 */

type RegisterHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type registerHandler struct {
	uc user.RegisterUsecase
	rr httplib.RequestReader
}

func (rh registerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := registerRequest{}
	err := rh.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := rh.uc.Register(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewRegisterHandler(uc user.RegisterUsecase, rr httplib.RequestReader) RegisterHandler {
	return &registerHandler{uc, rr}
}

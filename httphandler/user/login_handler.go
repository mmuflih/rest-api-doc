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
 * at: 2019-04-27 07:58
 */

type LoginHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type loginHandler struct {
	uc user.GetTokenUsecase
	rr httplib.RequestReader
}

func (lh loginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := registerRequest{}
	err := lh.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := lh.uc.Issue(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewLoginHandler(uc user.GetTokenUsecase, rr httplib.RequestReader) LoginHandler {
	return &loginHandler{uc, rr}
}

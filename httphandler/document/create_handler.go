package document

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/context/user"

	"github.com/mmuflih/rest-api-doc/context/document"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 10:08
 */

type CreateHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type createHandler struct {
	uc   document.CreateUsecase
	rr   httplib.RequestReader
	auth user.GetAuthUserUsecase
}

func (ch createHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := createRequest{}
	err := ch.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := ch.uc.Create(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewCreateHandler(uc document.CreateUsecase, rr httplib.RequestReader,
	auth user.GetAuthUserUsecase) CreateHandler {
	return &createHandler{uc, rr, auth}
}

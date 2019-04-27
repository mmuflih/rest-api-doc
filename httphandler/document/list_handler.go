package document

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/context/user"

	"github.com/mmuflih/rest-api-doc/context/document"
)

/*
 * Listd by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 10:08
 */

type ListHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type listHandler struct {
	uc   document.ListUsecase
	rr   httplib.RequestReader
	auth user.GetAuthUserUsecase
}

func (ch listHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := ch.rr.GetQuery(r, "pid")
	err, resp := ch.uc.List(id)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewListHandler(uc document.ListUsecase, rr httplib.RequestReader,
	auth user.GetAuthUserUsecase) ListHandler {
	return &listHandler{uc, rr, auth}
}

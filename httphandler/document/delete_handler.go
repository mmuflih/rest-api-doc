package document

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/context/user"

	"github.com/mmuflih/rest-api-doc/context/document"
)

/*
 * Deleted by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 10:08
 */

type DeleteHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type deleteHandler struct {
	uc   document.DeleteUsecase
	rr   httplib.RequestReader
	auth user.GetAuthUserUsecase
}

func (ch deleteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := ch.rr.GetRouteParam(r, "id")
	err, resp := ch.uc.Delete(id)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}

func NewDeleteHandler(uc document.DeleteUsecase, rr httplib.RequestReader,
	auth user.GetAuthUserUsecase) DeleteHandler {
	return &deleteHandler{uc, rr, auth}
}

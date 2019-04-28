package endpoint

import (
	"github.com/mmuflih/rest-api-doc/app"
	"github.com/mmuflih/rest-api-doc/domain/model"
	"github.com/mmuflih/rest-api-doc/domain/repository"
)

/*
 * Editd by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-28 05:42
 */

type EditUsecase interface {
	Edit(EditRequest) (error, interface{})
}

type EditRequest interface {
	GetDocumentID() string
	GetURL() string
	GetBody() string
	GetResponse() string
	GetAuthorization() string
	GetBodyOption() []model.BodyOption
}

type editUsecase struct {
	eRepo repository.EndpointRepository
}

func (cu editUsecase) Edit(req EditRequest) (error, interface{}) {
	err, e := cu.eRepo.FindByDocument(req.GetDocumentID())
	if err != nil {
		return app.NewReturnError("Data not found")
	}

	e.URL = req.GetURL()
	e.Body = req.GetBody()
	e.Response = req.GetResponse()
	e.BodyOption = req.GetBodyOption()
	e.Authorization = req.GetAuthorization()

	err = cu.eRepo.Update(e)
	if err != nil {
		return err, nil
	}
	return nil, e
}

func NewEditUsecase(eRepo repository.EndpointRepository) EditUsecase {
	return &editUsecase{eRepo}
}

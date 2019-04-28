package endpoint

import (
	"github.com/mmuflih/rest-api-doc/domain/model"
	"github.com/mmuflih/rest-api-doc/domain/repository"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-28 05:42
 */

type CreateUsecase interface {
	Create(CreateRequest) (error, interface{})
}

type CreateRequest interface {
	GetDocumentID() string
	GetURL() string
	GetMethod() string
	GetBody() string
	GetResponse() string
	GetAuthorization() string
	GetBodyOption() []model.BodyOption
}

type createUsecase struct {
	eRepo repository.EndpointRepository
}

func (cu createUsecase) Create(req CreateRequest) (error, interface{}) {
	e := model.NewEndpoint(req.GetDocumentID(), req.GetURL(), req.GetMethod(),
		req.GetBody(), req.GetResponse(), req.GetBodyOption(), req.GetAuthorization())
	err := cu.eRepo.Save(e)
	if err != nil {
		return err, nil
	}
	return nil, e
}

func NewCreateUsecase(eRepo repository.EndpointRepository) CreateUsecase {
	return &createUsecase{eRepo}
}

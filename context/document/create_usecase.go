package document

import (
	"github.com/mmuflih/rest-api-doc/domain/model"
	"github.com/mmuflih/rest-api-doc/domain/repository"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 09:43
 */

type CreateUsecase interface {
	Create(CreateRequest) (error, interface{})
}

type CreateRequest interface {
	GetID() string
	GetParentID() string
	GetName() string
}

type createUsecase struct {
	dRepo repository.DocumentRepository
}

func (cu createUsecase) Create(req CreateRequest) (error, interface{}) {
	doc := model.NewDocument(req.GetID(), req.GetParentID(), req.GetName())
	err, d := cu.dRepo.FindByID(req.GetParentID())
	if err == nil {
		d.Child = append(d.Child, doc)
		err = cu.dRepo.Update(d)
		return err, d
	}
	err = cu.dRepo.Save(doc)
	return err, doc
}

func NewCreateUsecase(dRepo repository.DocumentRepository) CreateUsecase {
	return &createUsecase{dRepo}
}

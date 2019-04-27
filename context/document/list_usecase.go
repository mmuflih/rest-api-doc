package document

import (
	"github.com/mmuflih/rest-api-doc/domain/repository"
)

/*
 * Listd by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 09:43
 */

type ListUsecase interface {
	List(parentID string) (error, interface{})
}

type listUsecase struct {
	dRepo repository.DocumentRepository
}

func (cu listUsecase) List(parentID string) (error, interface{}) {
	err, doc := cu.dRepo.List(parentID)
	return err, doc
}

func NewListUsecase(dRepo repository.DocumentRepository) ListUsecase {
	return &listUsecase{dRepo}
}

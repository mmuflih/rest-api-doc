package document

import (
	"github.com/mmuflih/rest-api-doc/app"
	"github.com/mmuflih/rest-api-doc/domain/repository"
)

/*
 * Deleted by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 09:43
 */

type DeleteUsecase interface {
	Delete(id string) (error, interface{})
}

type deleteUsecase struct {
	dRepo repository.DocumentRepository
}

func (cu deleteUsecase) Delete(id string) (error, interface{}) {
	err, doc := cu.dRepo.Find(id)
	if err != nil {
		return app.NewError("Data tidak ditemukan"), nil
	}
	err = cu.dRepo.Delete(id)
	return err, doc
}

func NewDeleteUsecase(dRepo repository.DocumentRepository) DeleteUsecase {
	return &deleteUsecase{dRepo}
}

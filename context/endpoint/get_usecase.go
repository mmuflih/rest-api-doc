package endpoint

import "github.com/mmuflih/rest-api-doc/domain/repository"

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-28 06:16
 */

type GetUsecase interface {
	Get(id string) (error, interface{})
}

type getUsecase struct {
	eRepo repository.EndpointRepository
}

func (gu getUsecase) Get(id string) (error, interface{}) {
	err, e := gu.eRepo.FindByDocument(id)
	if err != nil {
		return err, nil
	}

	return nil, e
}

func NewGetUsecase(eRepo repository.EndpointRepository) GetUsecase {
	return &getUsecase{eRepo}
}

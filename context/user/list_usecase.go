package user

import (
	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/domain/repository"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-27 11:15
 */

type ListUsecase interface {
	List(ListRequest) (error, ListResponse)
}

type ListRequest interface {
	GetQuery() string
	GetPage() int
	GetSize() int
}

type ListResponse interface {
}

type listUsecase struct {
	uRepo repository.UserRepository
}

func (lu listUsecase) List(req ListRequest) (error, ListResponse) {
	err, ls := lu.uRepo.FindAll(req.GetQuery(), req.GetPage(), req.GetSize())
	if err != nil {
		return err, nil
	}

	ul := []listResponse{}
	for _, u := range ls {
		lr := newListResponse(u)
		ul = append(ul, lr)
	}
	c := lu.uRepo.FindAllCount(req.GetQuery())

	return nil, httplib.NewDataPaginate(ul, c, req.GetPage(), req.GetSize())
}

func NewListUsecase(uRepo repository.UserRepository) ListUsecase {
	return &listUsecase{uRepo}
}

package user

import (
	"github.com/mmuflih/rest-api-doc/domain/model"
	"github.com/mmuflih/rest-api-doc/domain/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:29
**/

type AddUsecase interface {
	Add(AddRequest) (error, AddResponse)
}

type AddRequest interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetRole() string
	GetPhone() string
}

type AddResponse interface{}

type addUsecase struct {
	repo repository.UserRepository
}

func (au addUsecase) Add(req AddRequest) (error, AddResponse) {
	u := model.NewUser(req.GetEmail(), req.GetName(), req.GetPhone(), req.GetPassword(),
		req.GetRole())

	err := au.repo.Save(u)
	if err != nil {
		return err, nil
	}
	return err, u
}

func NewAddUsecase(repo repository.UserRepository) AddUsecase {
	return &addUsecase{repo}
}

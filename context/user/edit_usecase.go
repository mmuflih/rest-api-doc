package user

import (
	"github.com/mmuflih/rest-api-doc/domain/repository"
	"github.com/pkg/errors"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:37
**/

type EditUsecase interface {
	Edit(EditRequest) (error, EditResponse)
}

type EditRequest interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetRole() string
	GetPhone() string
}

type EditResponse interface{}

type editUsecase struct {
	repo repository.UserRepository
}

func (au editUsecase) Edit(req EditRequest) (error, EditResponse) {
	err, u := au.repo.Find(req.GetID())
	if err != nil {
		return errors.New("User not found"), nil
	}

	u.Name = req.GetName()
	u.Phone = req.GetPhone()
	u.Email = req.GetEmail()
	u.Password = req.GetPassword()
	u.Role = req.GetRole()

	err = au.repo.Update(u)
	if err != nil {
		return err, nil
	}
	return err, u
}

func NewEditUsecase(repo repository.UserRepository) EditUsecase {
	return &editUsecase{repo}
}

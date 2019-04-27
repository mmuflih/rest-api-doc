package user

import (
	"github.com/mmuflih/rest-api-doc/domain/model"
	"github.com/mmuflih/rest-api-doc/domain/repository"
	"github.com/mmuflih/rest-api-doc/lib"
	"github.com/mmuflih/rest-api-doc/role"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 08:58
 */

type GenerateAdminUsecase interface {
	Generate()
}

type generateAdminUsecase struct {
	uRepo repository.UserRepository
}

func (gauc generateAdminUsecase) Generate() {
	email := "backend@admin.com"
	pass := lib.GeneratePassword("@backend")
	err, _ := gauc.uRepo.FindByEmail(email)
	if err != nil {
		u := model.NewUser(email, "Backend Engineer", "085643332735", pass, role.BACKEND)
		u.Status = "approved"
		gauc.uRepo.Save(u)
	}
}

func NewGenerateAdminUsecase(uRepo repository.UserRepository) GenerateAdminUsecase {
	return &generateAdminUsecase{uRepo}
}

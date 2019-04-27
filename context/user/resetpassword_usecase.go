package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/rest-api-doc/domain/repository"
	"github.com/mmuflih/rest-api-doc/lib"
)

type ResetPasswordUsecase interface {
	Reset(email string) (error, ResetPasswordResponse)
}

type ResetPasswordResponse interface {
}

type resetPasswordUsecase struct {
	userRepo repository.UserRepository
	config   conf.Config
}

func NewResetPasswordUsecase(userRepo repository.UserRepository, config conf.Config) ResetPasswordUsecase {
	return &resetPasswordUsecase{userRepo, config}
}

func (rpu resetPasswordUsecase) Reset(email string) (error, ResetPasswordResponse) {
	err, u := rpu.userRepo.FindByEmail(email)
	if err != nil {
		return errors.New("User tidak ditemukan"), nil
	}

	pass := time.Now().Format("150405")
	u.Password = lib.GeneratePassword(pass)
	err = rpu.userRepo.Update(u)
	if err != nil {
		fmt.Println(err)
	}

	return nil, resetPasswordResponse{u}
}

type jsonParam struct {
	FirstName   string `json:"first_name" bson:"first_name"`
	Email       string `json:"email" bson:"email"`
	PasswordDec string `json:"password_dec" bson:"password_dec"`
}

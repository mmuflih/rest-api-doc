package user

import (
	"errors"

	"github.com/mmuflih/rest-api-doc/lib"

	"github.com/mmuflih/rest-api-doc/domain/repository"

	"github.com/mmuflih/envgo/conf"
	"golang.org/x/crypto/bcrypt"
)

type SetPasswordUsecase interface {
	SetPassword(SetPasswordRequest) (error, SetPasswordResponse)
}

type SetPasswordRequest interface {
	GetUserID() string
	GetOldPassword() string
	GetNewPassword() string
}
type SetPasswordResponse interface{}

type setPassword struct {
	userRepo repository.UserRepository
	config   conf.Config
}

func NewSetPasswordUsecase(userRepo repository.UserRepository, config conf.Config,
) SetPasswordUsecase {
	return &setPassword{userRepo, config}
}

func (this setPassword) SetPassword(r SetPasswordRequest) (error, SetPasswordResponse) {
	err, user := this.userRepo.Find(r.GetUserID())
	if err != nil {
		return err, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.GetOldPassword()))
	if err != nil {
		return errors.New("Password lama tidak cocok"), nil
	}

	user.Password = lib.GeneratePassword(r.GetNewPassword())
	err = this.userRepo.Update(user)
	if err != nil {
		return err, nil
	}
	return nil, user
}

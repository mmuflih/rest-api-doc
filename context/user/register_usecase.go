package user

import (
	"github.com/mmuflih/rest-api-doc/domain/model"
	"github.com/mmuflih/rest-api-doc/role"

	"github.com/mmuflih/rest-api-doc/domain/repository"
)

type RegisterUsecase interface {
	Register(RegisterRequest) (error, RegisterResponse)
}

type RegisterRequest interface {
	GetName() string
	GetEmail() string
	GetPhone() string
	GetPassword() string
}

type RegisterResponse interface{}

type registerUsecase struct {
	userRepo repository.UserRepository
	token    GetTokenUsecase
}

func NewRegisterUsecase(userRepo repository.UserRepository,
	token GetTokenUsecase) RegisterUsecase {
	return &registerUsecase{userRepo, token}
}

func (this registerUsecase) Register(req RegisterRequest) (error, RegisterResponse) {
	user := model.NewUser(req.GetEmail(), req.GetName(), req.GetPhone(), req.GetPassword(), role.FRONTEND)

	err := this.userRepo.Save(user)
	if err != nil {
		return err, nil
	}

	token := this.token.claimToken(user)
	return err, registerResponse{user, token}
}

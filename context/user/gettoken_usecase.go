package user

import (
	"errors"
	"time"

	"github.com/mmuflih/rest-api-doc/domain/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mmuflih/rest-api-doc/app"
	"github.com/mmuflih/rest-api-doc/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type GetTokenUsecase interface {
	Issue(GetTokenRequest) (error, interface{})
	claimToken(*model.User) model.AccessToken
}

type GetTokenRequest interface {
	GetEmail() string
	GetPassword() string
}

type getTokenUsecase struct {
	userRepo     repository.UserRepository
	signatureKey []byte
}

func NewGetTokenUsecase(userRepo repository.UserRepository,
	signatureKey []byte) GetTokenUsecase {
	return &getTokenUsecase{userRepo, signatureKey}
}

func (this *getTokenUsecase) Issue(req GetTokenRequest) (error, interface{}) {
	err, user := this.CheckUser(req.GetEmail())
	if err != nil {
		return err, nil
	}

	if user.Status != "approved" {
		return app.NewError("Contact your administrator"), nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))
	if err != nil {
		return app.NewError("Invalid username and password"), nil
	}
	at := this.claimToken(user)
	user.LastLogin = time.Now()
	err = this.userRepo.Update(user)
	if err != nil {
		return err, nil
	}

	return err, at
}

func (this *getTokenUsecase) CheckUser(email string) (error, *model.User) {
	err, user := this.userRepo.FindByEmail(email)
	if user.ID == "" {
		return errors.New("Username and password not match" + err.Error()), nil
	}
	return nil, user
}

func (this *getTokenUsecase) claimToken(u *model.User) model.AccessToken {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS512)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	expiredAt := time.Now().Add(time.Hour * (24 * 1500)).Unix()
	claims["user_id"] = u.ID
	claims["role"] = u.Role
	claims["exp"] = expiredAt

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(this.signatureKey)

	return model.AccessToken{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   expiredAt,
		Data:        u,
	}
}

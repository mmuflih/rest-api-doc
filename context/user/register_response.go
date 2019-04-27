package user

import "github.com/mmuflih/rest-api-doc/domain/model"

type registerResponse struct {
	*model.User
	AccessToken model.AccessToken `json:"access_token"`
}

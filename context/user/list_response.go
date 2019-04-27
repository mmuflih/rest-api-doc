package user

import "github.com/mmuflih/rest-api-doc/domain/model"

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-27 14:43
 */

type listResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	LastLogin string `json:"last_login"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func newListResponse(u *model.User) listResponse {
	return listResponse{
		ID:        u.ID.String(),
		Email:     u.Email,
		Name:      u.Name,
		Phone:     u.Phone,
		LastLogin: u.LastLogin.Format("2006-01-02 15:04:05"),
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

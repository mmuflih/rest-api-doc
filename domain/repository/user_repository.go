package repository

import (
	"github.com/mmuflih/rest-api-doc/domain/model"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:12
**/

type UserRepository interface {
	Save(u *model.User)error
	Update(u *model.User)error
	Find(id string) (error, *model.User)
	FindByEmail(email string) (error, *model.User)
	FindAll(q string, page, size int) (error, []*model.User)
	FindAllCount(q string) int
}

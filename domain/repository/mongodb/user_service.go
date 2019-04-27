package mongodb

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/mmuflih/rest-api-doc/domain/model"
	"github.com/mmuflih/rest-api-doc/domain/repository"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 07:02
 */

type userService struct {
	col *mgo.Collection
}

func (us userService) Save(u *model.User) error {
    return us.col.Insert(u)
}

func (us userService) Update(u *model.User) error {
    q := bson.M{"_id": u.ID}
    return us.col.Update(q, bson.M{"$set": u})
}

func (us userService) Find(id string) (error, *model.User) {
    u := new(model.User)
    err := us.col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
    return err, u
}

func (us userService) FindByEmail(email string) (error, *model.User) {
	u := new(model.User)
	err := us.col.Find(bson.M{"email": email}).One(&u)
	return err, u
}

func (us userService) FindAll(q string, page, size int) (error, []*model.User) {
	var list []*model.User
	err := us.col.Find(nil).All(&list)
	return err, list
}

func (us userService) FindAllCount(q string) int {
	panic("implement me")
}

func NewUserService(mdb *mgo.Database) repository.UserRepository {
	return &userService{mdb.C("users")}
}
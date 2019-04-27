package model

import (
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	Name      string        `json:"name" bson:"name"`
	Phone     string        `json:"phone" bson:"phone"`
	Password  string        `json:"-" bson:"password"`
	Role      string        `json:"role" bson:"role"`
	Status    string        `json:"status" bson:"status"`
	LastLogin time.Time     `json:"last_login" bson:"last_login"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func NewUser(email string, name string, phone string, password string, role string) *User {
	now := time.Now()
	user := new(User)
	user.ID = bson.NewObjectId()
	user.Email = email
	user.Name = name
	user.Phone = phone
	user.Password = password
	user.LastLogin = now
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Role = role
	user.Status = "pending"
	return user
}

func (u *User) GetFirstName() string {
	names := strings.Split(u.Name, " ")
	for _, name := range names {
		if len(name) > 2 {
			return name
		}
	}
	return names[0]
}

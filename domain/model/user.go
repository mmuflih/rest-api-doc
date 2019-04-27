package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type User struct {
	ID        bson.ObjectId    `json:"id" bson:"_id"`
	Email     string    `json:"email" bson:"email"`
	Name      string    `json:"name" bson:"name"`
	Phone     string    `json:"phone" bson:"phone"`
	Password  string    `json:"-" bson:"password"`
	Role      string    `json:"role" bson:"role"`
	LastLogin time.Time `json:"last_login" bson:"last_login"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
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
	return user
}

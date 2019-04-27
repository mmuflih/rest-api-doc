package model

import "github.com/globalsign/mgo/bson"

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 09:24
 */

type Endpoint struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

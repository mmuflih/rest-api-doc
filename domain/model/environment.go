package model

import "github.com/globalsign/mgo/bson"

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-28 13:37
 */

type Environment struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	ProjectID string        `json:"project_id" bson:"project_id"`
	Key       string        `json:"key" bson:"key"`
	Value     string        `json:"value" bson:"key"`
}

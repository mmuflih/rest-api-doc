package model

import (
	"time"

	"github.com/mmuflih/rest-api-doc/lib"

	"github.com/globalsign/mgo/bson"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 09:27
 */

type Document struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ProjectID string        `json:"project_id" bson:"project_id"`
	FolderID  string        `json:"folder_id" bson:"folder_id"`
	ParentID  string        `json:"parent_id" bson:"parent_id"`
	Method    string        `json:"method" bson:"method"`
	Name      string        `json:"name" bson:"name"`
	Child     []*Document   `json:"child" bson:"child"`
	DeletedAt lib.NullTime  `json:"deleted_at" bson:"deleted_at"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func NewDocument(folderID, parentID, name, method string) *Document {
	now := time.Now()
	id := bson.NewObjectId()
	return &Document{
		FolderID: folderID, ParentID: parentID, Name: name,
		CreatedAt: now, UpdatedAt: now, ID: id, Method: method}
}

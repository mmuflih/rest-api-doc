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
 * at: 2019-04-27 09:33
 */

type documentMongoService struct {
	col *mgo.Collection
}

func (fms documentMongoService) FindByID(id string) (error, *model.Document) {
	var docs *model.Document
	err := fms.col.Find(bson.M{"folder_id": id}).One(&docs)
	return err, docs
}

func (fms documentMongoService) Find(id string) (error, *model.Document) {
	var docs *model.Document
	err := fms.col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&docs)
	return err, docs
}

func (fms documentMongoService) List(parentID string) (error, []*model.Document) {
	var docs []*model.Document
	err := fms.col.Find(bson.M{"parent_id": parentID, "deleted_at.valid": false}).All(&docs)
	return err, docs
}

func (fms documentMongoService) Save(f *model.Document) error {
	return fms.col.Insert(f)
}

func (fms documentMongoService) Update(f *model.Document) error {
	q := bson.M{"_id": f.ID}
	return fms.col.Update(q, bson.M{"$set": f})
}

func (fms documentMongoService) Delete(id string) error {
	q := bson.M{"_id": bson.ObjectIdHex(id)}
	return fms.col.Remove(q)
}

func (fms documentMongoService) ListAll() (error, []*model.Document) {
	var docs []*model.Document
	err := fms.col.Find(bson.M{"deleted_at.valid": false}).All(&docs)
	return err, docs
}

func NewDocumentMongoService(mdb *mgo.Database) repository.DocumentRepository {
	return &documentMongoService{mdb.C("documents")}
}

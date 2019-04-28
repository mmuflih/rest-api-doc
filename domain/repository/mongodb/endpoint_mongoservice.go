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
 * at: 2019-04-28 05:36
 */

type endpointMongoService struct {
	col *mgo.Collection
}

func (ems endpointMongoService) Save(e *model.Endpoint) error {
	return ems.col.Insert(e)
}

func (ems endpointMongoService) Update(e *model.Endpoint) error {
	return ems.col.Update(bson.M{"_id": e.ID}, bson.M{"$set": e})
}

func (ems endpointMongoService) FindByDocument(id string) (error, *model.Endpoint) {
	e := new(model.Endpoint)
	err := ems.col.Find(bson.M{"document_id": id}).One(&e)
	return err, e
}

func (ems endpointMongoService) Delete(id string) error {
	return ems.col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

func NewEndpointMongoService(mdb *mgo.Database) repository.EndpointRepository {
	return &endpointMongoService{mdb.C("endpoints")}
}

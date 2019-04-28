package model

import "github.com/globalsign/mgo/bson"

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 09:24
 */

type Endpoint struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DocumentID    string        `json:"document_id" bson:"document_id"`
	Method        string        `json:"method" bson:"method"`
	URL           string        `json:"url" bson:"url"`
	Body          string        `json:"body" bson:"body"`
	Response      string        `json:"response" bson:"response"`
	BodyOption    []BodyOption  `json:"body_option" bson:"body_option"`
	Authorization string        `json:"authorization" bson:"authorization"`
}

func NewEndpoint(documentID, URL, method, body, response string,
	bodyOption []BodyOption, authorization string) *Endpoint {
	return &Endpoint{DocumentID: documentID, Method: method,
		URL: URL, Body: body, Response: response, BodyOption: bodyOption,
		Authorization: authorization,
	}
}

type BodyOption struct {
	Field       string `json:"field" bson:"field"`
	Required    bool   `json:"required" bson:"required"`
	DataType    string `json:"data_type" bson:"data_type"`
	Description string `json:"description" bson:"description"`
}

package endpoint

import "github.com/mmuflih/rest-api-doc/domain/model"

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-28 06:13
 */

type createRequest struct {
	DocumentID    string             `json:"document_id"`
	URL           string             `json:"url"`
	Method        string             `json:"method"`
	Body          string             `json:"body"`
	Response      string             `json:"response"`
	BodyOption    []model.BodyOption `json:"body_option"`
	Authorization string             `json:"authorization"`
}

func (cr createRequest) GetDocumentID() string {
	return cr.DocumentID
}
func (cr createRequest) GetURL() string {
	return cr.URL
}
func (cr createRequest) GetBody() string {
	return cr.Body
}
func (cr createRequest) GetResponse() string {
	return cr.Response
}
func (cr createRequest) GetBodyOption() []model.BodyOption {
	return cr.BodyOption
}
func (cr createRequest) GetAuthorization() string {
	return cr.Authorization
}
func (cr createRequest) GetMethod() string {
	return cr.Method
}

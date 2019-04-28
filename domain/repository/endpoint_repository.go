package repository

import "github.com/mmuflih/rest-api-doc/domain/model"

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-28 05:29
 */

type EndpointRepository interface {
	Save(e *model.Endpoint) error
	Update(e *model.Endpoint) error
	FindByDocument(id string) (error, *model.Endpoint)
	Delete(id string) error
}

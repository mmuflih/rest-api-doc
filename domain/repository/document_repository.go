package repository

import "github.com/mmuflih/rest-api-doc/domain/model"

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 09:31
 */

type DocumentRepository interface {
	Find(id string) (error, *model.Document)
	FindByID(id string) (error, *model.Document)
	Save(f *model.Document) error
	Update(f *model.Document) error
	Delete(id string) error
	ListAll() (error, []*model.Document)
	List(parentID string) (error, []*model.Document)
}

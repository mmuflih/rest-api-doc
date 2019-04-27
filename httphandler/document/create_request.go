package document

import (
	"strconv"
	"time"

	"github.com/gosimple/slug"
)

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 10:08
 */

type createRequest struct {
	ParentID string `json:"parent_id"`
	Name     string `json:"name"`
}

func (cr createRequest) GetID() string {
	return slug.Make(cr.Name) + strconv.Itoa(time.Now().Nanosecond())
}

func (cr createRequest) GetParentID() string {
	return cr.ParentID
}

func (cr createRequest) GetName() string {
	return cr.Name
}

package container

import (
	"github.com/mmuflih/rest-api-doc/domain/repository/mongodb"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func BuildRepositoryProvider(c *dig.Container) *dig.Container {
	if err := c.Provide(mongodb.NewUserService); err != nil {
		panic(err)
	}
	if err := c.Provide(mongodb.NewDocumentMongoService); err != nil {
		panic(err)
	}
	if err := c.Provide(mongodb.NewEndpointMongoService); err != nil {
		panic(err)
	}
	return c
}

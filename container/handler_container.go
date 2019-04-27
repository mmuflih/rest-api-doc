package container

import (
	"github.com/mmuflih/rest-api-doc/httphandler/extra"
	"github.com/mmuflih/rest-api-doc/httphandler/ping"
	"github.com/mmuflih/rest-api-doc/httphandler/user"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func BuildHandlerProvider(c *dig.Container) *dig.Container {
	if err := c.Provide(ping.NewPingHandler); err != nil {
		panic(err)
	}

	if err := c.Provide(extra.NewP404Handler); err != nil {
		panic(err)
	}

	if err := c.Provide(user.NewAddHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewEditHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewGetHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewListHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewLoginHandler); err != nil {
		panic(err)
	}
	if err := c.Provide(user.NewRegisterHandler); err != nil {
		panic(err)
	}

	return c
}

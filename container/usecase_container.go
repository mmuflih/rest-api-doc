package container

import (
	"github.com/mmuflih/go-di-arch/context/ping"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func BuildUseCaseProvider(c *dig.Container) *dig.Container {
	if err := c.Provide(ping.NewPingUsecase); err != nil {
		panic(err)
	}
	return c
}

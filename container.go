package main

import (
	"github.com/mmuflih/rest-api-doc/container"
	"go.uber.org/dig"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-05 10:10
**/

func BuildContainer() *dig.Container {
	c := dig.New()

	c = container.BuildConfigProvider(c)
	c = container.BuildRepositoryProvider(c)
	c = container.BuildUseCaseProvider(c)
	c = container.BuildHandlerProvider(c)

	return c
}

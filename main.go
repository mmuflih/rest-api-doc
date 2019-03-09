package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/mmuflih/go-di-arch/app"
	"github.com/mmuflih/go-httplib/httplib"
)

import "go.uber.org/dig"

var _ = dig.Name

func main() {
	myrole := make(map[string][]string)

	myrole[app.ADMIN] = []string{app.ADMIN}
	myrole[app.LEADER] = []string{app.LEADER, app.ADMIN}
	myrole[app.USER] = []string{app.USER, app.LEADER, app.ADMIN}

	httplib.InitJWTMiddlewareWithRole([]byte("Go-DI-arch"), jwt.SigningMethodHS512, myrole)

	c := BuildContainer()

	if err := c.Invoke(InvokeRoute); err != nil {
		panic(err)
	}

	if err := c.Provide(NewRoute); err != nil {
		panic(err)
	}

	if err := c.Invoke(func(s *ServerRoute) {
		s.Run()
	}); err != nil {
		fmt.Println(err)
	}
}

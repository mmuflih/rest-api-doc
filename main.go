package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/role"
)

import "go.uber.org/dig"

var _ = dig.Name

func main() {
	myrole := make(map[string][]string)

	myrole[role.BACKEND] = []string{role.BACKEND}
	myrole[role.FRONTEND] = []string{role.FRONTEND, role.BACKEND}

	httplib.InitJWTMiddlewareWithRole([]byte("12354p1d0c"), jwt.SigningMethodHS512, myrole)

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

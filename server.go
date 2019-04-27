package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/rest-api-doc/app"
	"github.com/mmuflih/rest-api-doc/context/user"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-05 10:59
**/

type ServerRoute struct {
	config  conf.Config
	handler http.Handler
	router  *mux.Router
	gauc    user.GenerateAdminUsecase
}

func NewRoute(c conf.Config, handler http.Handler, router *mux.Router,
	gauc user.GenerateAdminUsecase) *ServerRoute {
	for _, l := range c.GetStringSlice("env_label") {
		fmt.Println(l)
	}

	return &ServerRoute{c, handler, router, gauc}
}

func (s *ServerRoute) Run() {
	s.gauc.Generate()
	app.Logger("Application is running at ", time.Now().Format("2006-01-02 15:04:05.000"))
	app.Logger("Server listen on", s.config.GetString(`server.address`))
	log.Fatal(http.ListenAndServe(s.config.GetString(`server.address`), s.handler))
}

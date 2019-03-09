package main

import (
	"github.com/gorilla/mux"
	"github.com/mmuflih/go-di-arch/httphandler/extra"
	"github.com/mmuflih/go-di-arch/httphandler/ping"
	"net/http"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-09 20:49
**/
func InvokeRoute(route *mux.Router,
	pingH ping.PingHandler, p404H extra.P404Handler,
) {
	route.NotFoundHandler = http.HandlerFunc(p404H.Handle)
	/** api v1 route */
	apiV1 := route.PathPrefix("/api/v1").Subrouter()

	/** ping */
	pingRoute := apiV1.PathPrefix("/ping").Subrouter()
	pingRoute.HandleFunc("", pingH.Handle).Methods("GET")
}

package main

import (
	"net/http"

	"github.com/mmuflih/rest-api-doc/httphandler/endpoint"

	"github.com/mmuflih/rest-api-doc/role"

	jwt "github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/httphandler/document"

	"github.com/gorilla/mux"
	"github.com/mmuflih/rest-api-doc/httphandler/extra"
	"github.com/mmuflih/rest-api-doc/httphandler/ping"
	"github.com/mmuflih/rest-api-doc/httphandler/user"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-09 20:49
**/
func InvokeRoute(route *mux.Router,
	pingH ping.PingHandler, p404H extra.P404Handler, userAdd user.AddHandler,
	userEdit user.EditHandler, userGet user.GetHandler, userList user.ListHandler,
	registerH user.RegisterHandler, loginH user.LoginHandler,
	createDoc document.CreateHandler, deleteDoc document.DeleteHandler,
	listDoc document.ListHandler, createEndpiont endpoint.CreateHandler,
	editEndpoint endpoint.EditHandler, getEndpoint endpoint.EditHandler,
) {
	route.NotFoundHandler = http.HandlerFunc(p404H.Handle)
	/** api v1 route */
	apiV1 := route.PathPrefix("/api/v1").Subrouter()
	pingRoute := apiV1.PathPrefix("/ping").Subrouter()
	userRoute := apiV1.PathPrefix("/user").Subrouter()
	loginRoute := userRoute.PathPrefix("/login").Subrouter()
	documentRoute := apiV1.PathPrefix("/document").Subrouter()
	endpointRoute := apiV1.PathPrefix("/endpoint").Subrouter()

	/** ping */
	pingRoute.HandleFunc("", pingH.Handle).Methods("GET")

	/** user */
	userRoute.HandleFunc("/register", registerH.Handle).Methods("POST")
	userRoute.HandleFunc("", userAdd.Handle).Methods("POST")
	userRoute.HandleFunc("", userList.Handle).Methods("GET")
	userRoute.HandleFunc("/{id}", userEdit.Handle).Methods("PUT")
	userRoute.HandleFunc("/{id}", userGet.Handle).Methods("GET")

	/** login */
	loginRoute.HandleFunc("", loginH.Handle).Methods("POST")

	/** document */
	documentRoute.HandleFunc("", jwt.JWTMidWithRole(createDoc.Handle, role.BACKEND)).Methods("POST")
	documentRoute.HandleFunc("/{id}", jwt.JWTMidWithRole(deleteDoc.Handle, role.BACKEND)).Methods("DELETE")
	documentRoute.HandleFunc("", jwt.JWTMidWithRole(listDoc.Handle, role.FRONTEND)).Methods("GET")

	/** endpoint */
	endpointRoute.HandleFunc("", jwt.JWTMidWithRole(createEndpiont.Handle, role.BACKEND)).Methods("POST")
	endpointRoute.HandleFunc("", jwt.JWTMidWithRole(editEndpoint.Handle, role.BACKEND)).Methods("PUT")
	endpointRoute.HandleFunc("/{document_id}", jwt.JWTMidWithRole(getEndpoint.Handle, role.FRONTEND)).Methods("PUT")
}

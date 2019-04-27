package container

import (
	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"github.com/mmuflih/envgo/conf"
	"github.com/mmuflih/go-httplib/httplib"
	"github.com/mmuflih/rest-api-doc/config"
	"go.uber.org/dig"
	"net/http"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-06 19:47
**/

func BuildConfigProvider(c *dig.Container) *dig.Container {
	var err error

	if err = c.Provide(func() []byte {
		return []byte("12354p1d0c")
	}); err != nil {
		panic(err)
	}

	if err = c.Provide(func() conf.Config {
		return conf.NewConfig()
	}); err != nil {
		panic(err)
	}

	if err = c.Provide(func(c conf.Config) (error, *mgo.Database) {
		return config.NewMongoDB(c)
	}); err != nil {
		panic(err)
	}

	err = c.Provide(func() httplib.RequestReader {
		return httplib.NewMuxRequestReader()
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func() *mux.Router {
		return mux.NewRouter()
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func(c conf.Config, api *mux.Router) http.Handler {
		return config.InitCors(c, api)
	})
	if err != nil {
		panic(err)
	}

	return c
}

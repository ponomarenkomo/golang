package apiserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Start(config *Config) error {

	srv := httprouter.New()

	return http.ListenAndServe(config.BindAddr, srv)

}

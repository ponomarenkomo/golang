package apiserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func Start(config *Config) error {

	srv := httprouter.New()
	logrus.Info("Starting server!!!")

	return http.ListenAndServe(config.BindAddr, srv)

}

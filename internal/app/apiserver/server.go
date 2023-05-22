package apiserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *httprouter.Router
	logger *logrus.Logger
}

func NewServer() *server {
	s := &server{
		router: httprouter.New(),
		logger: logrus.New(),
	}

	return s

}

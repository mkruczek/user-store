package app

import (
	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
}

func newServer() *server {
	return &server{
		router: gin.Default(),
	}
}

func (s *server) start(arg string) {
	s.router.Run(arg) //todo handle error from Run()
}

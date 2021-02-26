package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mkruczek/user-store/config"
)

type server struct {
	router *gin.Engine
}

func newServer() *server {
	return &server{
		router: gin.Default(),
	}
}

func (s *server) start(cfg *config.Config) error {
	err := s.router.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		return err
	}
	return nil
}

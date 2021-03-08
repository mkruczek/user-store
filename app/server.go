package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mkruczek/user-store/config"
	"github.com/mkruczek/user-store/utils/logger"
)

type server struct {
	router *gin.Engine
	LOG    logger.Logger
}

func newServer(LOG logger.Logger) *server {
	return &server{
		router: gin.Default(),
		LOG:    LOG,
	}
}

func (s *server) start(cfg *config.Config) error {
	err := s.router.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		return err
	}
	return nil
}

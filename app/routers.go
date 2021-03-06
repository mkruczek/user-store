package app

import (
	"github.com/mkruczek/user-store/config"
	"github.com/mkruczek/user-store/controller/lifecheck"
	"github.com/mkruczek/user-store/controller/user"
)

func mapUrls(s *server, cfg *config.Config) {

	cu := user.NewUserController(cfg, s.LOG)

	s.router.GET("/ping", lifecheck.Check)

	s.router.POST("/user", cu.Create)
	s.router.GET("/user/:id", cu.GetById)
	s.router.GET("/user", cu.Search)
	s.router.PATCH("/user/:id", cu.PartialUpdate)
	s.router.PUT("/admin/user/:id", cu.FullUpdate) //todo secure
	s.router.DELETE("/user/:id", cu.Delete)
}

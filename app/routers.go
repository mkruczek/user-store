package app

import (
	"github.com/mkruczek/user-store/controller/lifecheck"
	"github.com/mkruczek/user-store/controller/user"
)

func mapUrls(s *server) {
	s.router.GET("/ping", lifecheck.Check) //not calling the function, only giving info what f need to be execution

	s.router.POST("/user", user.Create)
	s.router.GET("/user/:id", user.ById)
	s.router.GET("/user", user.Search)
	s.router.DELETE("/user/:id", user.Delete)
}

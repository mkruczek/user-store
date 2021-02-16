package app

import (
	"github.com/mkruczek/user-store/controller/lifecheck"
	"github.com/mkruczek/user-store/controller/user"
)

func mapUrls(s *server) {

	cu := user.NewUserController()

	s.router.GET("/ping", lifecheck.Check) //not calling the function, only giving info what f need to be execution

	s.router.POST("/user", cu.Create)
	s.router.GET("/user/:id", cu.ById)
	s.router.GET("/user", cu.Search)
	s.router.DELETE("/user/:id", cu.Delete)
}

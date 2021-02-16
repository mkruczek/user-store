package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mkruczek/user-store/domain/user"
	userService "github.com/mkruczek/user-store/service/user"
	"github.com/mkruczek/user-store/utils/errors"
	"net/http"
)

type Controller struct {
	userService *userService.Service
}

func NewUserController() *Controller {
	return &Controller{userService: userService.NewUserService()}
}

func (c *Controller) Create(g *gin.Context) {

	var newUser user.DTO
	if err := g.ShouldBindJSON(&newUser); err != nil {
		g.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	created, err := c.userService.CreateUser(newUser)
	if err != nil {
		g.JSON(http.StatusBadRequest, errors.NewInternalServerError(err.Error()))
		return
	}

	g.JSON(http.StatusCreated, created)
}
func (c *Controller) Search(g *gin.Context) {
	g.String(http.StatusNotImplemented, "implement me!")
}
func (c *Controller) ById(g *gin.Context) {
	g.String(http.StatusNotImplemented, "implement me!")
}

func (c *Controller) Delete(g *gin.Context) {
	g.String(http.StatusNotImplemented, "implement me!")
}

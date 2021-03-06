package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mkruczek/user-store/config"
	"github.com/mkruczek/user-store/domain/user"
	userService "github.com/mkruczek/user-store/domain/user/service/user"
	"github.com/mkruczek/user-store/utils/errors"
	"github.com/mkruczek/user-store/utils/logger"
	"net/http"
	"strings"
)

type Controller struct {
	userService *userService.Service
	LOG         logger.Logger
}

func NewUserController(cfg *config.Config, log logger.Logger) *Controller {
	return &Controller{
		userService: userService.NewUserService(cfg),
		LOG:         log,
	}
}

func (c *Controller) Create(g *gin.Context) {

	var newUser user.PrivateView
	if err := g.ShouldBindJSON(&newUser); err != nil {
		c.LOG.Warnf("problem with parsing request : %s", err.Error())
		g.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	created, err := c.userService.Create(newUser)
	if err != nil {
		c.LOG.Warnf("problem during created new User : %v", err)
		g.JSON(err.Status, err)
		return
	}
	c.LOG.Infof("creating new user : %v", created)
	g.JSON(http.StatusCreated, created)
}

func (c *Controller) PartialUpdate(g *gin.Context) {
	var newValue user.UpdateDTO
	if err := g.ShouldBindJSON(&newValue); err != nil {
		g.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	id := g.Param("id")
	updated, err := c.userService.PartialUpdate(id, newValue)
	if err != nil {
		g.JSON(err.Status, err)
		return
	}

	g.JSON(http.StatusCreated, updated)
}

func (c *Controller) FullUpdate(g *gin.Context) {
	var newValue user.PrivateView
	if err := g.ShouldBindJSON(&newValue); err != nil {
		g.JSON(http.StatusBadRequest, errors.NewBadRequestError(err.Error()))
		return
	}

	id := g.Param("id")
	updated, err := c.userService.FullUpdate(id, newValue)
	if err != nil {
		g.JSON(err.Status, err)
		return
	}

	g.JSON(http.StatusCreated, updated)
}

func (c *Controller) Search(g *gin.Context) {
	values := getSearchValues(g)

	result, err := c.userService.Search(g.GetHeader("X-Public") == "true", values)
	if err != nil {
		g.JSON(err.Status, err)
		return
	}

	g.JSON(http.StatusOK, result)
}

func (c *Controller) GetById(g *gin.Context) {
	id := g.Param("id")
	result, err := c.userService.GetById(g.GetHeader("X-Public") == "true", id)
	if err != nil {
		g.JSON(err.Status, err)
		return
	}

	g.JSON(http.StatusOK, result)
}

func (c *Controller) Delete(g *gin.Context) {
	id := g.Param("id")
	err := c.userService.Delete(id)
	if err != nil {
		g.JSON(err.Status, err)
		return
	}

	g.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func getSearchValues(g *gin.Context) map[string][]string {
	values := make(map[string][]string)
	ids := g.Query("id")
	if ids != "" {
		values["id"] = strings.Split(ids, ",")
	}
	firstNames := g.Query("first_name")
	if firstNames != "" {
		values["first_name"] = strings.Split(firstNames, ",")
	}
	lastNames := g.Query("last_name")
	if lastNames != "" {
		values["last_name"] = strings.Split(lastNames, ",")
	}
	emails := g.Query("email")
	if emails != "" {
		values["email"] = strings.Split(emails, ",")
	}

	//todo handle createDate == < >
	return values
}

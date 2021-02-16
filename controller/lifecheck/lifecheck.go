package lifecheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

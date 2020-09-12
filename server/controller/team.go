package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type _Team struct{}

var TeamController *_Team

func (_ *_Team) Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}

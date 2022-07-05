package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) FlutterAccess(c *gin.Context) {
	if verify(c, s) {
		res := map[string]interface{}{
			"data": "Hello world!",
		}
		c.JSON(http.StatusOK, res)
	}
}

func (s *server) HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

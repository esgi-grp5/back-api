package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Access without OAuth */

func (s *server) HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

/* Access with OAuth */

func (s *server) FlutterAccess(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Hello world!",
	}
	c.JSON(http.StatusOK, res)
}

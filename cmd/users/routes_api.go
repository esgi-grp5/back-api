package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FlutterAccess godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Unauthorized 401 {object} map[string]interface{}
// @Router /flutter [get]
func (s *server) FlutterAccess(c *gin.Context) {
	if verify(c, s) {
		res := map[string]interface{}{
			"data": "Flutter have accesses",
		}
		c.JSON(http.StatusOK, res)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (s *server) HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

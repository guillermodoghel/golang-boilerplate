package server

import "github.com/gin-gonic/gin"

type PingController interface {
	Ping(c *gin.Context)
}

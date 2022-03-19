package server

import "github.com/gin-gonic/gin"

type PingController interface  {
	Ping(c *gin.Context)
}

type TimeController interface  {
	Time(c *gin.Context)
}
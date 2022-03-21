package server

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"guillermodoghel/golang-boilerplate/src/rest"
)

type Router struct {
	Engine *gin.Engine
}

func (r Router) SetupRoutes(pingController PingController) {
	r.Engine.GET("/ping", pingController.Ping)
}

func NewRouter() *Router {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(RequestIdMiddleware())
	router.NoRoute(rest.NewStatusNotFound)

	return &Router{
		Engine: router,
	}
}

func (r Router) Run() error {
	return r.Engine.Run()
}

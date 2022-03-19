package server

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"guillermodoghel/golang-boilerplate/internal/rest"
)

const defaultPort = "8080"

type Router struct {
	Engine *gin.Engine
}

func (r Router) SetupRoutes(pingController PingController, timeController TimeController) {
	r.Engine.GET("/ping", pingController.Ping)
	r.Engine.GET("/time", timeController.Time)
}

func (r Router) Run() error {
	return r.Engine.Run()
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

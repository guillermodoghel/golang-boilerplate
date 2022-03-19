package server

import (
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"guillermodoghel/golang-boilerplate/internal/logger"
	"guillermodoghel/golang-boilerplate/internal/rest"
)

const defaultPort = ":8080"

type Router struct {
	Engine *gin.Engine
}

func (r Router) SetupRoutes(pingController PingController, timeController TimeController) {
	r.Engine.POST("/ping", pingController.Ping)
	r.Engine.GET("/time", timeController.Time)
}

func (r Router) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		logger.GetLogger().Warnf("WARNING: 'PORT' unspecified. starting server on default port '%s'\n", defaultPort)
		port = defaultPort
	}
	return r.Engine.Run(port)
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

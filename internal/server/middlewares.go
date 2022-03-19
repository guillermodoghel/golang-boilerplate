package server

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const XRequestId = "X-Request-Id"

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4().String()
		defer func() {
			c.Writer.Header().Set(XRequestId, uuid)
			return
		}()
		c.Next()
	}
}
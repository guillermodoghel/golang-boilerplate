package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func NewStatusInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, map[string]string{
		"message": http.StatusText(http.StatusInternalServerError),
	})
}


func StatusNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, map[string]string{
		"message": http.StatusText(http.StatusNotFound),
	})
}

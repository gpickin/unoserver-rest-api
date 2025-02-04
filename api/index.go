package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	log.Printf("UnoServer is up")
	c.String(http.StatusOK, "UnoServer is up")
	return
}
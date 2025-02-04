package api

import (
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
}

func ListenAndServe(addr string) {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Routes
	router.GET("/", Index)
	router.GET("/healthcheck", Index)
	router.POST("/request", RequestHandler)

	if addr == "" {
		addr = ":2004"
	}

	pm := endless.NewServer(addr, router)
	pm.BeforeBegin = func(add string) {
		log.Printf("Server is running at %s", addr)
		log.Printf("Server is running pid is %d", syscall.Getpid())
	}

	pm.ListenAndServe()
}

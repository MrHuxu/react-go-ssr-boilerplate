package server

import (
	"github.com/gin-gonic/gin"
)

func init() {
	s := &server{
		Engine: gin.Default(),
	}
	s.initRenderFunc()
	s.registerRoutes()

	DefaultServer = s
}

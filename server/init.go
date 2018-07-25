package server

import (
	"github.com/gin-gonic/gin"
)

func init() {
	s := &server{
		Engine: gin.Default(),
	}
	s.registerRoutes()
	s.initReact()

	DefaultServer = s
}

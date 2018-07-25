package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var DefaultServer Server

type Server interface {
	Run()
}

type server struct {
	*gin.Engine

	react string
}

func (s *server) Run() {
	s.Engine.Run(fmt.Sprintf(":%d", port))
}

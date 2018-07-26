package server

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

var DefaultServer Server

type Server interface {
	Run()
}

type server struct {
	*gin.Engine

	jsVM         *goja.Runtime
	jsRenderFunc func(goja.FunctionCall) goja.Value
}

func (s *server) Run() {
	s.Engine.Run(fmt.Sprintf(":%d", port))
}

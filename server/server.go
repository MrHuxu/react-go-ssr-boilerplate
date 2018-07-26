package server

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/react-go-ssr-boilerplate/server/conf"
)

var DefaultServer Server

func init() {
	s := &server{
		Engine: gin.Default(),
	}
	s.registerRoutes()

	DefaultServer = s
}

type Server interface {
	Run()
}

type server struct {
	*gin.Engine

	jsVM     *goja.Runtime
	jsRender func(goja.FunctionCall) goja.Value
}

func (s *server) Run() {
	s.Engine.Run(fmt.Sprintf(":%d", conf.Conf.Web.Port))
}

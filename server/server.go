package server

import (
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/react-go-ssr-boilerplate/server/conf"
	"github.com/MrHuxu/react-go-ssr-boilerplate/server/react"
)

var DefaultServer Server

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

func (s *server) renderHTMLString(url string, data interface{}) string {
	urlVal, _ := react.VM.RunString(fmt.Sprintf("'%s'", url))

	bytes, _ := json.Marshal(data)
	dataVal, _ := react.VM.RunString(fmt.Sprintf("'%s'", string(bytes)))

	return react.Render(goja.FunctionCall{Arguments: []goja.Value{urlVal, dataVal}}).String()
}

func init() {
	s := &server{
		Engine: gin.Default(),
	}
	s.LoadHTMLGlob(conf.Conf.Web.TemplatesPath)
	s.registerRoutes()

	DefaultServer = s
}

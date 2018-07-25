package server

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.Engine.GET("/", func(ctx *gin.Context) {
		vm := goja.New()
		v, err := vm.RunString(s.react)
		if err != nil {
			panic(err)
		}

		fn := v.Export().(map[string]interface{})["genHtmlString"].(func(goja.FunctionCall) goja.Value)
		url, _ := vm.RunString(fmt.Sprintf("\"%s\"", ctx.Request.URL.String()))

		ctx.Writer.Write([]byte(fn(goja.FunctionCall{Arguments: []goja.Value{url}}).String()))
		return
	})

	s.Engine.GET("/test", func(ctx *gin.Context) {
		vm := goja.New()
		v, err := vm.RunString(s.react)
		if err != nil {
			panic(err)
		}

		fn := v.Export().(map[string]interface{})["genHtmlString"].(func(goja.FunctionCall) goja.Value)
		url, _ := vm.RunString(fmt.Sprintf("\"%s\"", ctx.Request.URL.String()))

		ctx.Writer.Write([]byte(fn(goja.FunctionCall{Arguments: []goja.Value{url}}).String()))
		return
	})
}

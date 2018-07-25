package server

import (
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.Engine.GET("/test", func(ctx *gin.Context) {
		vm := goja.New()
		_, err := vm.RunString(s.react)
		if err != nil {
			panic(err)
		}

		var fn func(string) string
		vm.ExportTo(vm.Get("genHtmlString"), &fn)

		ctx.Writer.Write([]byte(fn(ctx.Request.URL.String())))
		return
	})
}

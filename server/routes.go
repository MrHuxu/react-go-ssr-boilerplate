package server

import (
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.Engine.GET("/test", func(ctx *gin.Context) {
		vm := goja.New()
		v, err := vm.RunString(s.react)
		if err != nil {
			panic(err)
		}

		ctx.Writer.Write([]byte(v.Export().(map[string]interface{})["html"].(string)))
		return
	})
}

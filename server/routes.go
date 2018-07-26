package server

import (
	"encoding/json"
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
		url, _ := vm.RunString(fmt.Sprintf("'%s'", ctx.Request.URL.String()))

		bytes, _ := json.Marshal(map[string]string{"home": "test home in store hehe"})
		data, _ := vm.RunString(fmt.Sprintf("'%s'", string(bytes)))

		ctx.Writer.Write([]byte(fn(goja.FunctionCall{Arguments: []goja.Value{url, data}}).String()))
		return
	})

	s.Engine.GET("/test", func(ctx *gin.Context) {
		vm := goja.New()
		v, err := vm.RunString(s.react)
		if err != nil {
			panic(err)
		}

		_ = v.Export().(map[string]interface{})["genHtmlString"].(func(goja.FunctionCall) goja.Value)
		_, _ = vm.RunString(fmt.Sprintf("'%s'", ctx.Request.URL.String()))

		bytes, _ := json.Marshal(map[string]string{"home": "test home in store"})
		data, _ := vm.RunString(string(bytes))
		fmt.Println("test")
		fmt.Println(data)

		ctx.Writer.Write([]byte("test"))
		// ctx.Writer.Write([]byte(fn(goja.FunctionCall{Arguments: []goja.Value{url, data}}).String()))
		return
	})
}

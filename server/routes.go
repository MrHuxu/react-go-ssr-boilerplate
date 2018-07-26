package server

import (
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/react-go-ssr-boilerplate/server/react"
)

func (s *server) registerRoutes() {
	s.Engine.GET("/", func(ctx *gin.Context) {
		url, _ := react.VM.RunString(fmt.Sprintf("'%s'", ctx.Request.URL.String()))

		bytes, _ := json.Marshal(map[string]interface{}{
			"home": map[string]interface{}{
				"titles": []string{"test title"},
				"infos":  map[string]interface{}{"test title": 1},
			},
		})
		data, _ := react.VM.RunString(fmt.Sprintf("'%s'", string(bytes)))

		ctx.Writer.Write([]byte(react.Render(goja.FunctionCall{Arguments: []goja.Value{url, data}}).String()))
		return
	})

	s.Engine.GET("/test", func(ctx *gin.Context) {
		url, _ := react.VM.RunString(fmt.Sprintf("'%s'", ctx.Request.URL.String()))

		bytes, _ := json.Marshal(map[string]interface{}{
			"test": map[string]interface{}{
				"list":  []string{"test"},
				"infos": map[string]int{"test": 2},
			},
		})
		data, _ := react.VM.RunString(fmt.Sprintf("'%s'", string(bytes)))

		ctx.Writer.Write([]byte(react.Render(goja.FunctionCall{Arguments: []goja.Value{url, data}}).String()))
		return
	})
}

package server

import (
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.Engine.GET("/", func(ctx *gin.Context) {
		url, _ := s.jsVM.RunString(fmt.Sprintf("'%s'", ctx.Request.URL.String()))

		bytes, _ := json.Marshal(map[string]interface{}{
			"home": map[string]interface{}{
				"titles": []string{"test title"},
				"infos":  map[string]interface{}{"test title": 1},
			},
		})
		data, _ := s.jsVM.RunString(fmt.Sprintf("'%s'", string(bytes)))

		ctx.Writer.Write([]byte(s.jsRenderFunc(goja.FunctionCall{Arguments: []goja.Value{url, data}}).String()))
		return
	})

	s.Engine.GET("/test", func(ctx *gin.Context) {
		url, _ := s.jsVM.RunString(fmt.Sprintf("'%s'", ctx.Request.URL.String()))

		bytes, _ := json.Marshal(map[string]interface{}{
			"test": map[string]interface{}{
				"list":  []string{"test"},
				"infos": map[string]int{"test": 2},
			},
		})
		data, _ := s.jsVM.RunString(fmt.Sprintf("'%s'", string(bytes)))

		ctx.Writer.Write([]byte(s.jsRenderFunc(goja.FunctionCall{Arguments: []goja.Value{url, data}}).String()))
		return
	})
}

package server

import (
	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte(s.renderHTMLString(
			ctx.Request.URL.String(),
			map[string]interface{}{
				"home": map[string]interface{}{
					"titles": []string{"test title"},
					"infos":  map[string]interface{}{"test title": 1},
				},
			},
		)))
	})

	s.Engine.GET("/test", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte(s.renderHTMLString(
			ctx.Request.URL.String(),
			map[string]interface{}{
				"test": map[string]interface{}{
					"list":  []string{"test"},
					"infos": map[string]int{"test": 2},
				},
			},
		)))
	})
}

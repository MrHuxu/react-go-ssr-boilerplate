package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.Engine.GET("/", func(ctx *gin.Context) {
		println("huhu")
		ctx.Set("res", map[string]interface{}{
			"data": map[string]interface{}{
				"home": map[string]interface{}{
					"titles": []string{"test title"},
					"infos":  map[string]interface{}{"test title": 1},
				},
			},
		})
	})

	s.Engine.GET("/test", func(ctx *gin.Context) {
		println("hihi")
		ctx.Set("res", map[string]interface{}{
			"data": map[string]interface{}{
				"test": map[string]interface{}{
					"list":  []string{"test"},
					"infos": map[string]int{"test": 2},
				},
			},
		})
	})

	s.Engine.GET("/api/test", func(ctx *gin.Context) {
		println("hehe")
		ctx.JSON(http.StatusOK, gin.H{
			"test": "this is a test",
		})
	})
}

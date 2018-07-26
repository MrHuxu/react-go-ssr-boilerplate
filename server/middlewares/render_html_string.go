package middlewares

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/react-go-ssr-boilerplate/server/react"
)

func RenderHTMLString() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if strings.HasPrefix(ctx.Request.URL.String(), "/api") {
			return
		}

		res, _ := ctx.Get("res")
		pageInfo := convertResToPageInfo(ctx.Request.URL.String(), res)
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H(pageInfo))
	}
}

func convertResToPageInfo(url string, res interface{}) map[string]interface{} {
	var resMap map[string]interface{}
	if m, ok := res.(map[string]interface{}); ok {
		resMap = m
	}

	var meta, title, data interface{}
	if d, ok := resMap["meta"]; ok {
		meta = d
	}
	if d, ok := resMap["title"]; ok {
		title = d
	}
	if d, ok := resMap["data"]; ok {
		data = d
	}

	urlVal, _ := react.VM.RunString(fmt.Sprintf("'%s'", url))
	bytes, _ := json.Marshal(data)
	dataVal, _ := react.VM.RunString(fmt.Sprintf("'%s'", string(bytes)))
	bodyStr := react.Render(goja.FunctionCall{Arguments: []goja.Value{urlVal, dataVal}}).String()

	return map[string]interface{}{
		"meta":  meta,
		"title": title,
		"body":  template.HTML(bodyStr),
	}
}

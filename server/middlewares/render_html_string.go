package middlewares

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
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

		var pageInfo map[string]interface{}
		defer func() {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H(pageInfo))
		}()

		status := ctx.Writer.Status()
		if status != 200 {
			urlVal, _ := react.VM.RunString(fmt.Sprintf("'/%d'", status))
			pageInfo = map[string]interface{}{
				"meta":  fmt.Sprintf("%d Error", status),
				"title": "Error Happened!",
				"body":  template.HTML(react.Render(goja.FunctionCall{Arguments: []goja.Value{urlVal}}).String()),
			}
			return
		}

		res, _ := ctx.Get("res")
		pageInfo = convertResToPageInfo(ctx.Request.URL, res)
		return
	}
}

func convertResToPageInfo(url *url.URL, res interface{}) map[string]interface{} {
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

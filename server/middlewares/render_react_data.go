package middlewares

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/react-go-ssr-boilerplate/server/react"
)

// RenderReactData renders the react data to given template
func RenderReactData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if strings.HasPrefix(ctx.Request.URL.String(), "/api") {
			return
		}

		pageInfo := getPageInfo(ctx)
		ctx.HTML(ctx.Writer.Status(), "index.tmpl", gin.H(pageInfo))
	}
}

func getPageInfo(ctx *gin.Context) map[string]interface{} {
	status := ctx.Writer.Status()
	if status != 200 {
		return map[string]interface{}{
			"meta":  fmt.Sprintf("%d Error", status),
			"title": "Error Happened!",
			"body":  template.HTML(react.Render(fmt.Sprintf("/%d", status), nil)),
		}
	}

	res, _ := ctx.Get("res")
	return getPageInfoFromRes(ctx.Request.URL.String(), res)
}

func getPageInfoFromRes(url string, res interface{}) map[string]interface{} {
	var resMap map[string]interface{}
	if m, ok := res.(map[string]interface{}); ok {
		resMap = m
	}

	pageInfo := make(map[string]interface{})
	if m, ok := resMap["meta"]; ok {
		pageInfo["meta"] = m
	}
	if t, ok := resMap["title"]; ok {
		pageInfo["title"] = t
	}
	if d, ok := resMap["data"]; ok {
		pageInfo["body"] = template.HTML(react.Render(url, d))
	} else {
		pageInfo["body"] = template.HTML(react.Render(url, nil))
	}
	return pageInfo
}

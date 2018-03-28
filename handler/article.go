package handler

import (
	"gin-api-example/handler/v1"
	"gin-api-example/handler/v2"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
}

// 重新分派路由
func (a *ArticleHandler) Index(c *gin.Context) {

	version, _ := c.Get("version")
	if version == "v1" {
		handler := v1.ArticleHandler{}
		handler.Index(c)
	} else {
		handler := v2.ArticleHandler{}
		handler.Index(c)
	}

}

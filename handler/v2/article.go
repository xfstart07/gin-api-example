package v2

import (
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
}

func (a *ArticleHandler) Index(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "article v2",
	})

}

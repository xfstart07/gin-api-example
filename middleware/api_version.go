package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ApiVersion() gin.HandlerFunc {
	return func(c *gin.Context) {

		accept := c.GetHeader("Accept")
		fmt.Println("accept", accept)
		// example.com 一般是公司的域名
		if accept == "application/vnd.example.com.v1+json" {
			c.Set("version", "v1")
		} else {
			c.Set("version", "v2")
		}

		// request
		c.Next()
		// respose
	}
}

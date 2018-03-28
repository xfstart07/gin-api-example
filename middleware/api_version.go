package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ApiVersion() gin.HandlerFunc {
	return func(c *gin.Context) {

		accept := c.GetHeader("Accept")
		fmt.Println("accept", accept)
		if accept == "application/vnd.bjywkd.v1+json" {
			c.Set("version", "v1")
		} else {
			c.Set("version", "v2")
		}

		// request
		c.Next()
		// respose
	}
}

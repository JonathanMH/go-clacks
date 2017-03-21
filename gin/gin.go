package goClacks

import "gopkg.in/gin-gonic/gin.v1"

func Terrify() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("X-Clacks-Overhead", "GNU Terry Pratchett")
		c.Next()
	}
}

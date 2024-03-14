package middlewares

import "github.com/gin-gonic/gin"

// func Authenticate() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if !(c.Request.Header.Get("Token") == "true") {
// 			c.AbortWithStatusJSON(500, gin.H{
// 				"message": "Token Not Present",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }

func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Token not present",
		})
		return
	}
	c.Next()
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key", "Value")
	c.Next()
}

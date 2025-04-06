package middleware

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Custom logger logic
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Next()
	}
}

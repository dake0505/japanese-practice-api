package middleware

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered from panic: %v", r)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()

		c.Next()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.UnknownError, gin.H{
					"error": e.Error(),
				}))
			}
			c.Abort()
		}
	}
}

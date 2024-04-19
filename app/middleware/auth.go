package middleware

import (
	"context"
	"gin-gonic-api/app/firebase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")
		app := firebase.InitFirebase()

		client, err := app.Auth(c)

		token, err := client.VerifyIDToken(context.Background(), idToken)

		println(token, "tokentokentoken")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// Token is valid, set the user ID in the context
		c.Set("userID", token.UID)
		c.Next()
	}
}

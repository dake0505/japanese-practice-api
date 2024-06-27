package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-gonic-api/app/firebase"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FirebaseCustomTokenRequest struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type FirebaseIDTokenResponse struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func GetIDTokenFromCustomToken(customToken string) (string, error) {
	apiKey := "AIzaSyBgjNxcZqSH_xUnb7buABmupvXO5_5XXxs"
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s", apiKey)
	reqBody := FirebaseCustomTokenRequest{
		Token:             customToken,
		ReturnSecureToken: true,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var idTokenResponse FirebaseIDTokenResponse
	if err := json.Unmarshal(body, &idTokenResponse); err != nil {
		return "", err
	}
	return idTokenResponse.IDToken, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		customToken := c.GetHeader("Authorization")
		log.Printf("token %v", customToken)
		app := firebase.InitFirebase()
		client, err := app.Auth(c.Request.Context())
		idToken, err := GetIDTokenFromCustomToken(customToken)
		token, err := client.VerifyIDToken(c.Request.Context(), idToken)
		log.Printf("error ============================= %v", err)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set("userAuthID", token.UID)

		userRecord, err := client.GetUser(c.Request.Context(), token.UID)
		c.Set("userRecord", userRecord)
		c.Next()
	}
}

package authmiddleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// MiddlewareAuth adalah middleware untuk validasi token dari auth service
func MiddlewareAuth(c *gin.Context) {
	// Ambil URL dari environment variable
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server error", "message": "AUTH_SERVICE_URL is not set"})
		return
	}

	// Cek header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Authorization header missing"})
		return
	}

	// Validasi token dengan auth service
	authRes, err := validateToken(authServiceURL, authHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": err.Error()})
		return
	}

	// Set user data ke context untuk dipakai di handler lain
	c.Set("USER_ID", authRes.Data.ID)
	c.Set("FULL_NAME", authRes.Data.FullName)
	c.Set("EMAIL", authRes.Data.Email)
	c.Set("ACCOUNT_ACTIVE", authRes.Data.AccountActive)

	// Lanjutkan request
	c.Next()
}

// validateToken mengirim request ke auth service untuk validasi token
func validateToken(authServiceURL, token string) (*AuthResponse, error) {
	// Buat request ke auth/validate
	req, err := http.NewRequest("GET", authServiceURL, nil)
	if err != nil {
		return nil, errors.New("failed to create request")
	}
	req.Header.Set("Authorization", token)

	// Kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("failed to connect to auth service")
	}
	defer resp.Body.Close()

	// Parse response JSON
	var authRes AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authRes); err != nil {
		return nil, errors.New("failed to parse auth response")
	}

	// Jika token tidak valid
	if !authRes.IsSuccess || authRes.Data == nil {
		return nil, errors.New(authRes.Message)
	}

	return &authRes, nil
}

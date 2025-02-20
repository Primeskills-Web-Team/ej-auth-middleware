# ej-auth-middleware 🚀  
Middleware autentikasi JWT untuk layanan microservice berbasis Go.

## 🔧 Instalasi  
Tambahkan library ini ke proyek Go:  
```sh
go get github.com/Primeskills-Web-Team/ej-auth-middleware@latest
```

## 🚀 Penggunaan  
```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Primeskills-Web-Team/ej-auth-middleware/authmiddleware"
)

func main() {
	r := gin.Default()

	// Middleware untuk seluruh route
	r.Use(authmiddleware.MiddlewareAuth)

	r.GET("/profile", func(c *gin.Context) {
		userID, _ := c.Get("USER_ID")
		fullName, _ := c.Get("FULL_NAME")

		c.JSON(200, gin.H{
			"user_id":   userID,
			"full_name": fullName,
		})
	})

	// Middleware untuk group route tertentu
	adminGroup := r.Group("/admin")
	adminGroup.Use(authmiddleware.MiddlewareAuth)
	{
		adminGroup.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to Admin Dashboard"})
		})
	}

	r.Run(":8080")
}
```

## ⚙️ Cara Kerja  
1. Middleware membaca `Authorization` header.
2. Token divalidasi ke endpoint `auth/validate`.
3. Jika token valid, data pengguna disimpan di `gin.Context`.
4. Jika token tidak valid, request dibatalkan dengan status dari auth service.

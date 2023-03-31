package middleware

import (
	"fmt"
	"net/http"

	token "github.com/MarkoVasilic/Accommodation-booking-platform/airplane_tickets_app/server/tokens"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("Authorization")
		if ClientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func CheckIsRoleRegular() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("Authorization")
		if ClientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, _ := token.ValidateToken(ClientToken)
		if claims.Role != "REGULAR" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not allowed for that role"})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func CheckIsRoleAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("Authorization")
		if ClientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, _ := token.ValidateToken(ClientToken)
		if claims.Role != "ADMIN" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not allowed for that role"})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		fmt.Println("A")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		fmt.Println("B")

		c.Next()
	}
}

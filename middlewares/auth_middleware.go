package middlewares

import (
	"net/http"
	"os"

	"github.com/ank809/Event-Management-golang~/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, "error: Missing cookie")
				c.Abort()
				return
			} else {
				c.JSON(http.StatusBadRequest, err)
				c.Abort()
				return
			}
		}
		claims := &models.Claims{}

		if err := godotenv.Load(".env"); err != nil {
			c.JSON(http.StatusForbidden, err)
			c.Abort()
			return
		}
		JWT_KEY := []byte(os.Getenv("JWT_KEY"))
		token, err := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, err)
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, err)
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusBadRequest, "Invalid token")
			c.Abort()
			return
		}
		c.Set("user", claims)
		c.Next()

	}
}

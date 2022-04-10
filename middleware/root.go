package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-swagger-tutorial/pkg/jwt"
	"net/http"
)

func AuthorizeJWT(jwt jwt.JWTService) gin.HandlerFunc {

	return func(c *gin.Context) {

		var (
			isValidToken bool
			err          error
		)

		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BearerSchema):]

		if _, isValidToken, err = jwt.DecodeJWT(tokenString); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !isValidToken {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

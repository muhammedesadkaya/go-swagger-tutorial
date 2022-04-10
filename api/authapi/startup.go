package authapi

import (
	. "github.com/go-swagger-tutorial/pkg/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-swagger-tutorial/pkg/jwt"
)

// Auth godoc
// @Summary      Generate Token
// @Description  Generate token
// @Tags         Auths
// @Param 		 request body AuthRequest true "request"
// @Accept       json
// @Produce      json
// @Failure      400  {object} HTTPError
// @Failure      404  {object} HTTPError
// @Failure      500  {object} HTTPError
// @Router       /auth [post]
func Auth(r *gin.Engine, jwt jwt.JWTService) {

	r.POST("/auth", func(c *gin.Context) {

		var (
			requestBody *AuthRequest
			tokenString string
			err         error
		)

		if err = c.BindJSON(&requestBody); err != nil {
			NewError(c, http.StatusBadRequest, err)
			return
		}

		if !isValidUser(requestBody.Username, requestBody.Password) {
			c.JSON(http.StatusNotFound, gin.H{"token": ""})
			return
		}

		if tokenString, err = jwt.GenerateJWT(requestBody.Username); err != nil {
			NewError(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})
}

func isValidUser(username, password string) bool {

	var (
		dummyUsername = "muhammedesadkaya"
		dummyPassword = "123456"
	)

	return dummyUsername == username && dummyPassword == password
}

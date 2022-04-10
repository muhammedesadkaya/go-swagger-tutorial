package userapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
}

// GetUsers User godoc
// @Summary      Get All Users
// @Description  Get All Users
// @Security     BearerAuth
// @Tags         Users
// @Accept       json
// @Produce      json
// @Failure      400  {object}  httperror.HTTPError
// @Failure      404  {object}  httperror.HTTPError
// @Failure      500  {object}  httperror.HTTPError
// @Router       /users [get]
func GetUsers(r *gin.Engine) {

	r.GET("/users", func(c *gin.Context) {

		dummyData := User{
			FullName: "Muhammed Esad Kaya",
			PhoneNumber: "05000000000",
		}

		c.JSON(http.StatusOK, dummyData)

	})
}

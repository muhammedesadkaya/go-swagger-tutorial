package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-swagger-tutorial/api/authapi"
	"github.com/go-swagger-tutorial/api/userapi"
	. "github.com/go-swagger-tutorial/configs"
	_ "github.com/go-swagger-tutorial/docs"
	"github.com/go-swagger-tutorial/middleware"
	"github.com/go-swagger-tutorial/pkg/jwt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Golang Swagger Tutorial
// @version         1.0
// @description     Golang Swagger Github Tutorial
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  muhammedesadkaya@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:1453

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	var jwtService = jwt.NewJWTService(SecretKey, ExpireTimeMinute)

	authapi.Auth(router, jwtService)

	router.Use(middleware.AuthorizeJWT(jwtService))
	{
		userapi.GetUsers(router)
	}

	router.Run(":1453")
}

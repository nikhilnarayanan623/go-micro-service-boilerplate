package routes

import (
	"api-gateway/pkg/api/handler/interfaces"
	handlerInterface "api-gateway/pkg/api/handler/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(
	studentHandler interfaces.StudentHandler,
	authHandler handlerInterface.AuthHandler,
) http.Handler {

	router := gin.New()

	router.Use(gin.Logger())

	// group with api version
	api := router.Group("/api/v1")

	auth := api.Group("auth")
	{ // for all routes under auth

		RegisterAuthRoutes(auth, authHandler)
	}

	// student := api.Group("/student")
	{
		// group student routes
		// register student handler with student routes
		// RegisterStudentRoutes(student, studentHandler)
	}

	return router
}

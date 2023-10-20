package routes

import (
	"api-gateway/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(studentHandler interfaces.StudentHandler) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())

	{
		// group student routes
		student := router.Group("/student")
		// register student handler with student routes
		RegisterStudentRoutes(student, studentHandler)
	}

	return router
}

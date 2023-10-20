package routes

import (
	handlerInterface "api-gateway/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(student *gin.RouterGroup, studentHandler handlerInterface.StudentHandler) {

	student.POST("", studentHandler.Create)

}

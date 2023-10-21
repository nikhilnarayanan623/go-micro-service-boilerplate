package routes

import (
	handlerInterfaces "api-gateway/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(student *gin.RouterGroup, studentHandler handlerInterfaces.StudentHandler) {

	student.POST("", studentHandler.Create)

}

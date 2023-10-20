package routes

import (
	"api-gateway/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func RegisterStudentRoutes(student *gin.RouterGroup, studentHandler interfaces.StudentHandler) {

	student.POST("", studentHandler.Create)

}

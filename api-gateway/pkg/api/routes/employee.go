package routes

import (
	handlerInterfaces "api-gateway/pkg/api/handler/interfaces"
	"api-gateway/pkg/api/middleware/interfaces"

	"github.com/gin-gonic/gin"
)

func RegisterEmployeeRoutes(
	employees *gin.RouterGroup,
	middleware interfaces.Middleware,
	studentHandler handlerInterfaces.EmployeeHandler,
) {

	employees.POST("", studentHandler.Create)

}

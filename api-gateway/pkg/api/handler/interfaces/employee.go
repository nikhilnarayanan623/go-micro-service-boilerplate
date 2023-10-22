package interfaces

import (
	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
	Create(ctx *gin.Context)
}

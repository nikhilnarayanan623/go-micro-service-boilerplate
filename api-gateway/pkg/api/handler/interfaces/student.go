package interfaces

import (
	"github.com/gin-gonic/gin"
)

type StudentHandler interface {
	Create(ctx *gin.Context)
}

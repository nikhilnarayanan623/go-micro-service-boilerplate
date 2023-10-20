package handler

import (
	"api-gateway/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

type studentHandler struct {
}

func NewStudentHandler() interfaces.StudentHandler {
	return &studentHandler{}
}

func (s *studentHandler) Create(ctx *gin.Context) {

}

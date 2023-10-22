package handler

import (
	"api-gateway/pkg/api/handler/interfaces"
	"api-gateway/pkg/api/handler/request"
	"api-gateway/pkg/api/handler/response"
	clientInterfaces "api-gateway/pkg/client/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type employeeHandler struct {
	client clientInterfaces.EmployeeServiceClient
}

func NewEmployeeHandler(client clientInterfaces.EmployeeServiceClient) interfaces.EmployeeHandler {
	return &employeeHandler{
		client: client,
	}
}

// @Summary		Create Employees
// @Description	API for User To Create Random Employees Details
// @Security		BearerAuth
// @Id				Create Employees
// @Tags			Employees
// @Param			inputs	body	request.Employee{}	true	"Create Employees Details"
// @Router			/employees [post]
// @Success		200	{object}	response.Response{data=[]response.Employee}	"Successfully Employees Created"
// @Failure		500	{object}	response.Response{}							"Internal Server Error"
func (s *employeeHandler) Create(ctx *gin.Context) {

	var body request.Employee

	// bind request  values to body and validate
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(http.StatusBadRequest, BindErrorMessage, err, body)
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	employees, err := s.client.Create(ctx, body)

	if err != nil {
		response := response.ErrorResponse(http.StatusInternalServerError, "internal server error", err, nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := response.SuccessResponse(http.StatusOK, "Successfully employees created", employees)
	ctx.JSON(http.StatusOK, response)
}

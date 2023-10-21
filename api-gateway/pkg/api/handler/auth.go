package handler

import (
	"api-gateway/pkg/api/handler/interfaces"
	"api-gateway/pkg/api/handler/request"
	"api-gateway/pkg/api/handler/response"
	clientInterfaces "api-gateway/pkg/client/interfaces"
	"api-gateway/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	client clientInterfaces.AuthServiceClient
}

func NewAuthHandler(client clientInterfaces.AuthServiceClient) interfaces.AuthHandler {

	return &authHandler{
		client: client,
	}
}

func (a *authHandler) SignUp(ctx *gin.Context) {

	// create request body which have binding validation tags
	var body request.SignUp

	// use gin binding which will also check the validation and return error if any.
	if err := ctx.ShouldBindJSON(&body); err != nil {

		response := response.ErrorResponse(http.StatusBadRequest, BindErrorMessage, err, body)
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	// call the client to do the sign up
	res, err := a.client.SignUp(ctx, body)

	if err != nil {
		// log the error
		log.Println(err)

		// get the status code from the error
		statusCode := utils.GetHTTPStatusCodeFromGRPCError(err)
		var message string

		// crate a message according to the status code
		switch statusCode {
		case http.StatusConflict:
			message = "user already exist with given details"
		default:
			message = "internal server error"
		}

		response := response.ErrorResponse(statusCode, message, err, nil)
		ctx.JSON(statusCode, response)
		return
	}

	response := response.SuccessResponse(http.StatusCreated, "successfully sign up completed", res)

	ctx.JSON(http.StatusCreated, response)
}

func (a *authHandler) SignIn(ctx *gin.Context) {

	// create request body which have binding validation tags
	var body request.SignIn

	// use gin binding which will also check the validation and return error if any.
	if err := ctx.ShouldBindJSON(&body); err != nil {

		response := response.ErrorResponse(http.StatusBadRequest, BindErrorMessage, err, body)
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	// call the client to do the sign in
	res, err := a.client.SignIn(ctx, body)

	if err != nil {
		// log the error
		log.Println(err)

		// get the status code from the error
		statusCode := utils.GetHTTPStatusCodeFromGRPCError(err)
		var message string

		// crate a message according to the status code
		switch statusCode {
		case http.StatusNotFound:
			message = "user not found with given details"
		case http.StatusUnauthorized:
			message = "user password doesn't match"
		default:
			message = "internal server error"
		}

		response := response.ErrorResponse(statusCode, message, err, nil)
		ctx.JSON(statusCode, response)
		return
	}

	response := response.SuccessResponse(http.StatusCreated, "successfully sign in completed", res)

	ctx.JSON(http.StatusCreated, response)
}

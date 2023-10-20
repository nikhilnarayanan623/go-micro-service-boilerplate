package routes

import (
	handlerInterface "api-gateway/pkg/api/handler/interfaces"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(auth *gin.RouterGroup, authHandler handlerInterface.AuthHandler) {

	signUp := auth.Group("/sign-up")
	{
		signUp.POST("", authHandler.SignUp)
	}

	signIn := auth.Group("/sign-in")
	{

	}
	_ = signIn
}

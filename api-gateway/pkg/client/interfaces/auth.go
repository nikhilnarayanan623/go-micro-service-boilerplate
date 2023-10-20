package interfaces

import (
	"api-gateway/pkg/api/handler/request"
	"api-gateway/pkg/api/handler/response"
	"context"
)

type AuthServiceClient interface {
	SignUp(ctx context.Context, signUpDetails request.SignUp) (response.SignUp, error)
}

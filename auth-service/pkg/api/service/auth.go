package service

import (
	"auth-service/pkg/domain"
	"auth-service/pkg/pb"
	"auth-service/pkg/usecase"
	"auth-service/pkg/usecase/interfaces"
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServiceServer struct {
	authUseCase interfaces.AuthUseCase

	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(authUseCase interfaces.AuthUseCase) pb.AuthServiceServer {

	return &authServiceServer{
		authUseCase: authUseCase,
	}
}

func (a *authServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	// create user with request use details.
	user := domain.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	user, err := a.authUseCase.SignUp(user)

	if err != nil {
		// log the error
		log.Println(err)

		var (
			statusCode codes.Code
			message    string
		)
		// check the error and according to the error set status code and message
		switch {
		case errors.Is(err, usecase.ErrAlreadyExist):
			statusCode = codes.AlreadyExists
			message = "user already exist with given details"
		default:
			statusCode = codes.Internal
			message = "internal server error"
		}
		return nil, status.Error(statusCode, message)
	}

	return &pb.SignUpResponse{
		UserId: user.ID26,
	}, nil
}

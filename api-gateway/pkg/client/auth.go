package client

import (
	"api-gateway/pkg/api/handler/request"
	"api-gateway/pkg/api/handler/response"
	"api-gateway/pkg/client/interfaces"
	"api-gateway/pkg/config"
	"auth-service/pkg/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authServiceClient struct {
	client pb.AuthServiceClient
}

// This client is abstraction over the actual client
func NewAuthServiceClient(cfg config.Config) (interfaces.AuthServiceClient, error) {

	// create the auth service address
	addr := fmt.Sprintf("%s:%s", cfg.AuthServiceHost, cfg.AuthServicePort)

	// create a grpc client connection to auth service url
	cc, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create a grpc client connection for auth service : %w", err)
	}

	// create new auth service client with the grpc connection
	client := pb.NewAuthServiceClient(cc)

	// return our abstracted client with grpc auth service client
	return &authServiceClient{
		client: client,
	}, nil
}

func (a *authServiceClient) SignUp(ctx context.Context, signUpDetails request.SignUp) (response.SignUp, error) {

	// create a proto generated request with our sign up details
	req := &pb.SignUpRequest{
		Email:    signUpDetails.Email,
		Password: signUpDetails.Password,
	}

	// call to the actual client sign method invoke grpc call
	res, err := a.client.SignUp(ctx, req)

	// if any error just simply return and handler the error on handler
	if err != nil {
		return response.SignUp{}, err
	}

	// return the user id with our response struct
	return response.SignUp{
		UserID: res.GetUserId(),
	}, nil
}

package client

import (
	"api-gateway/pkg/api/handler/request"
	"api-gateway/pkg/api/handler/response"
	"api-gateway/pkg/client/interfaces"
	"api-gateway/pkg/config"
	"api-gateway/pkg/pb"
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type employeeServiceClient struct {
	client pb.EmployeeServiceClient
}

// This client is abstraction over the actual client
func NewEmployeeServiceClient(cfg config.Config) (interfaces.EmployeeServiceClient, error) {

	// create the employee service address
	addr := fmt.Sprintf("%s:%s", cfg.EmployeeServiceHost, cfg.EmployeeServicePort)

	// create a grpc client connection to auth service url
	cc, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create a grpc client connection for auth service : %w", err)
	}

	// create new employee service client with the grpc connection
	client := pb.NewEmployeeServiceClient(cc)

	// return our abstracted client with grpc employee service client
	return &employeeServiceClient{
		client: client,
	}, nil
}

func (e *employeeServiceClient) Create(ctx context.Context, empReq request.Employee) ([]response.Employee, error) {

	// create request
	req := &pb.CreateRequest{
		Count: int32(empReq.Count),
	}

	// call client to get stream client
	stream, err := e.client.Create(ctx, req)

	if err != nil {
		return nil, err
	}

	var employees []response.Employee

	for {
		fmt.Println("receiving...")
		res, err := stream.Recv()
		fmt.Println("received: ", res, err)
		if err != nil {
			// check error is EOF
			if err == io.EOF {
				// stream completed
				break
			}
			return nil, err
		}

		// copy the employee details to response employee
		employee := response.Employee{
			ID:    res.GetId(),
			Name:  res.GetName(),
			Email: res.GetEmail(),
			Age:   int(res.GetAge()),
			Role:  res.GetRole(),
		}
		// append the employee to employees
		employees = append(employees, employee)
	}

	return employees, nil
}

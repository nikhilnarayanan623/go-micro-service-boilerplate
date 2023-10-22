package service

import (
	"context"
	"employee-service/pkg/domain"
	"employee-service/pkg/pb"
	"employee-service/pkg/usecase/interfaces"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type employeeServiceServer struct {
	employeeUseCase interfaces.EmployeeUseCase

	pb.UnimplementedEmployeeServiceServer
}

func NewAuthServiceServer(employeeUseCase interfaces.EmployeeUseCase) pb.EmployeeServiceServer {

	return &employeeServiceServer{
		employeeUseCase: employeeUseCase,
	}
}

func (e *employeeServiceServer) Create(req *pb.CreateRequest, stream pb.EmployeeService_CreateServer) error {

	// create a channel to get employee from usecase
	employeeChan := make(chan domain.Employee)

	// create a context with cancel to cancel notify the usecase once the stream completed or need to exit because of error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run the employee usecase in different goroutines to get the employee details concurrently
	go e.employeeUseCase.Create(ctx, int(req.GetCount()), employeeChan)

	// create a max wait duration for each employee creation
	maxWaitDuration := time.Second * 10

	for i := 1; i <= int(req.GetCount()); i++ {
		select {
		case emp := <-employeeChan:

			err := stream.Send(&pb.Employee{
				Id:    emp.ID,
				Name:  emp.Name,
				Age:   int32(emp.Age),
				Email: emp.Email,
				Role:  emp.Role,
			})

			if err != nil {
				log.Println("error while sending employee as stream: %w", err)
				return status.Error(codes.Internal, fmt.Sprintf("failed to send stream: %v", err))
			}
		case <-time.After(maxWaitDuration):
			// if the max time duration of employee creation exceeded return.
			log.Println("employee creation on server time exceeded")
			return status.Error(codes.DeadlineExceeded, "employee creation on serve time exceed")
		}
	}

	return nil
}

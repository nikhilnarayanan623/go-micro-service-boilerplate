package api

import (
	"employee-service/pkg/config"
	"employee-service/pkg/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	lis  net.Listener
	gsr  *grpc.Server
	port string
}

func NewServerGRPC(cfg config.Config, srv pb.EmployeeServiceServer) (*Server, error) {

	addr := fmt.Sprintf("%s:%s", cfg.EmployeeServiceHost, cfg.EmployeeServicePort)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	gsr := grpc.NewServer()

	pb.RegisterEmployeeServiceServer(gsr, srv)

	return &Server{
		lis:  lis,
		gsr:  gsr,
		port: cfg.EmployeeServicePort,
	}, err
}

func (c *Server) Start() error {
	log.Println("Employee service listening on port: ", c.port)
	return c.gsr.Serve(c.lis)
}

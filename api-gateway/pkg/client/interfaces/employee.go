package interfaces

import (
	"api-gateway/pkg/api/handler/request"
	"api-gateway/pkg/api/handler/response"
	"context"
)

type EmployeeServiceClient interface {
	Create(ctx context.Context, req request.Employee) ([]response.Employee, error)
}

package interfaces

import (
	"context"
	"employee-service/pkg/domain"
)

type EmployeeUseCase interface {
	Create(ctx context.Context, count int, employeeChan chan<- domain.Employee)
}

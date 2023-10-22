package usecase

import (
	"context"
	"employee-service/pkg/domain"
	"employee-service/pkg/service/random"
	"employee-service/pkg/usecase/interfaces"
)

type employeeUseCase struct {
	randomGen random.RandomGenerator
}

func NewEmployeeUseCase(randomGen random.RandomGenerator) interfaces.EmployeeUseCase {

	return &employeeUseCase{
		randomGen: randomGen,
	}
}

// To create n number of employee's details
func (e *employeeUseCase) Create(ctx context.Context, count int, employeeChan chan<- domain.Employee) {

	mailSet := make(map[string]struct{}, count)

	for i := 1; i <= count; i++ {

		emp := e.randomGen.CreateEmployee()
		// check any unique filed is again randomly generated(ex: email)
		if _, ok := mailSet[emp.Email]; ok {
			// decrease the i and create random employee again
			continue
		}

		// everything ok the send the employee to employee channel
		select {
		case <-ctx.Done(): // check the context done for exit
			return
		default:
			employeeChan <- emp
		}
	}
}

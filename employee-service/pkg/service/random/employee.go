package random

import (
	"employee-service/pkg/domain"
	"employee-service/pkg/utils"
	"math/rand"
	"strings"
	"sync"
)

const (
	employeeMinAge = 18
	employeeMaxAge = 50
)

var (
	mu    sync.RWMutex
	roles = [...]string{"Software Developer",
		"Network Administrator",
		"Database Administrator (DBA)",
		"Cybersecurity Analyst",
		"System Administrator",
		"Data Scientist",
		"Cloud Solutions Architect",
		"IT Project Manager",
		"Quality Assurance/Analyst Engineer(QA)",
		"DevOps Engineer",
	}
)

func (r *randomGenerator) CreateEmployee() domain.Employee {

	// select a random name from names
	name := r.getRandomName()

	return domain.Employee{
		ID:    utils.GenerateUUID(),
		Name:  name,
		Email: createEmail(name),
		Age:   utils.GetIntBetween(employeeMinAge, employeeMaxAge),
		Role:  getRandomRole(),
	}

}

func createEmail(name string) string {

	// change the name to lowercase
	name = strings.ToLower(name)

	numChar := []byte("123456789")

	extraLetters := 4

	// cover the name to byte slice as email
	email := []byte(name)

	for i := 1; i <= extraLetters; i++ {
		// take a random number character from numChan and append to email
		email = append(email, numChar[rand.Intn(len(numChar))])
	}

	mailEnd := []byte("@email.com")
	// add the email ending to the email
	email = append(email, mailEnd...)

	// convert the mail slice to string and return
	return string(email)
}

func getRandomRole() string {
	// do a read lock for roles
	mu.RLock()
	defer mu.RUnlock()
	//	return a random role
	return roles[utils.GetRandomIndex(len(roles))]
}

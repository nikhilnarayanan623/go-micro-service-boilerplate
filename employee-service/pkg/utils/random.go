package utils

import (
	"math/rand"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.NewString()
}

// To get random index. out will within(0 to len-1)
func GetRandomIndex(len int) int {
	return rand.Intn(len)
}

// To get an int between start and end(start and end is inclusive)
func GetIntBetween(start, end int) int {

	diff := (end - start) + 1

	return start + (rand.Intn(diff))
}

package random

import "github.com/google/uuid"

// RandomNumber returns a random number.
func RandomNumber() string {
	return uuid.New().String()
}

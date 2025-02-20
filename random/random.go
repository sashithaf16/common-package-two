package random

import "github.com/google/uuid"

// RandomNumber returns a random number.
func RandomNumber() string {
	str := uuid.New().String()
	return str + "HELLO" + "UPDATED"
}

package random

import (
	"math/rand"
)

func RandomNumber(n int) int {
	// rand.Seed(n * int(time.Now().UnixNano()) depraceted
	random := rand.Intn(n)
	return random
}

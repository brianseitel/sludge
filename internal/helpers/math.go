package helpers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Fuzzy fuzzes a number
func Fuzzy(num int) int {
	fuzz := rand.Intn(10)

	switch {
	case fuzz < 3:
		num--
	case fuzz > 6:
		num++
	}

	if num > 1 {
		return num
	}

	return 1
}

package randx

import (
	"crypto/rand"
	"math/rand"
	"time"
	"fmt"
)

func InitToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func InitSeed() {
	rand.Seed(time.Now().UnixNano())
}

func IntBetween(low, high int) int {
	if low == high {
		return low
	}
	if high < low {
		high, low = low, high
	}
	// [low, high)
	return low + rand.Intn(high-low)
}

func Int64Between(low, high int64) int64 {
	if low == high {
		return low
	}
	if high < low {
		high, low = low, high
	}
	// [low, high)
	return low + rand.Int63n(high-low)
}

func Float64Between(low, high float64) float64 {
	if low == high {
		return low
	}
	if high < low {
		high, low = low, high
	}
	// [low, high)
	return low + rand.Float64()*(high-low)
}

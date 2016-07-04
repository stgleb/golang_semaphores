package boston_court

import (
	"time"
	"math/rand"
)

func randSleep(low, high int) {
	const N = 10 // time scale
	time.Sleep(randMillisecond(low * N, high * N))
}

func randMillisecond(low, high int) time.Duration {
	return time.Duration(low + rand.Intn(high - low)) * time.Millisecond
}


package retry

import (
	"time"
)

func BackoffLinear(waitBetween time.Duration) BackoffFunc {
	_ = "STUB: not implemented"
	return *new(BackoffFunc)
}

func jitterUp(duration time.Duration, jitter float64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func exponentBase2(a uint) uint64 { _ = "STUB: not implemented"; return 0 }

func BackoffLinearWithJitter(waitBetween time.Duration, jitterFraction float64) BackoffFunc {
	_ = "STUB: not implemented"
	return *new(BackoffFunc)
}

func BackoffExponential(scalar time.Duration) BackoffFunc {
	_ = "STUB: not implemented"
	return *new(BackoffFunc)
}

func BackoffExponentialWithJitter(scalar time.Duration, jitterFraction float64) BackoffFunc {
	_ = "STUB: not implemented"
	return *new(BackoffFunc)
}

func BackoffExponentialWithJitterBounded(scalar time.Duration, jitterFrac float64, maxBound time.Duration) BackoffFunc {
	_ = "STUB: not implemented"
	return *new(BackoffFunc)
}

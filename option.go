package repeat

import (
	"math/rand"
	"time"
)

// OptFunc function for runner configuration.
type OptFunc func(*Runner)

// WithAttemps returns the OptFunc function
// to set the number of attempts.
func WithAttemps(attempts int) OptFunc {
	return func(r *Runner) {
		r.attempts = attempts
	}
}

// WithDelay returns the OptFunc function
// to set the delay between attempts.
func WithDelay(d time.Duration) OptFunc {
	return func(r *Runner) {
		r.delay = func() time.Duration { return d }
	}
}

// WithBackoffDelay returns the OptFunc function
// to set the delay between attempts
// using the exponential delay algorithm.
func WithBackoffDelay(startDelay, maxDelay, jitterDelay time.Duration) OptFunc {
	return func(r *Runner) {
		baseDelay := startDelay / 2
		r.delay = func() time.Duration {
			baseDelay *= 2
			if baseDelay > maxDelay {
				baseDelay = maxDelay
			}

			baseDelay += jitter(jitterDelay)

			return baseDelay
		}
	}
}

// jitter returns a random value within [-span, span) range
func jitter(span time.Duration) time.Duration {
	return time.Duration(float64(span) * (2.0*rand.Float64() - 1.0))
}

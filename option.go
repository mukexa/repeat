// Copyright © 2017 Artem Feschenko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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

// WithDelayFunc returns the OptFunc function
// to set the delay between attempts using the custom function.
func WithDelayFunc(fn func() time.Duration) OptFunc {
	return func(r *Runner) {
		r.delay = fn
	}
}

// WithBackoffDelay returns the OptFunc function
// to set the delay between attempts
// using the exponential delay algorithm.
func WithBackoffDelay(startDelay, maxDelay, jitterDelay time.Duration) OptFunc {
	return func(r *Runner) {
		baseDelay := startDelay
		if baseDelay > 1 {
			baseDelay /= 2
		}
		r.delay = func() time.Duration {
			baseDelay *= 2
			if baseDelay > maxDelay {
				baseDelay = maxDelay
			}
			return baseDelay + jitter(jitterDelay)
		}
	}
}

// jitter returns a random value within [-span, span] range
func jitter(span time.Duration) time.Duration {
	return span * time.Duration(rand.Int31n(3)-1)
}

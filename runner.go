// Copyright © 2017 Artem Feschenko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repeat

import (
	"context"
	"time"
)

// DefaultAttempts default number of attempts.
const DefaultAttempts = -1

// Reasons for execution stops.
const (
	AttemptsReached   Reason = "attempts is reached"
	ContextDoneSignal Reason = "received completion signal from the context"
)

// Reason the execution stop reason.
type Reason string

func (r Reason) String() string {
	return string(r)
}

// Runner main execution controller.
type Runner struct {
	attempts int
	counter  int
	ctx      context.Context
	delay    func() time.Duration
	reason   Reason
}

// NewRunner returns new Runner with default values.
func NewRunner(ctx context.Context, opts ...OptFunc) *Runner {
	r := &Runner{
		attempts: DefaultAttempts,
		counter:  0,
		ctx:      ctx,
		delay:    func() time.Duration { return 0 },
	}

	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Reason returns reason for execution stops.
func (r *Runner) Reason() Reason {
	return r.reason
}

// Count returns count of execution attempts.
func (r *Runner) Count() int {
	return r.counter
}

// Next returns true if next execution is needed.
// In case false, runner contains stopping reason.
func (r *Runner) Next() bool {
	if r.attempts > 0 {
		if r.counter >= r.attempts {
			r.reason = AttemptsReached
			return false
		}
	}

	// guarantee the first execute and avoid a delay before this.
	if r.counter == 0 {
		return true
	}

	select {
	case <-r.ctx.Done():
		r.reason = ContextDoneSignal
		return false
	case <-time.After(r.delay()):
	}

	return true
}

// Execute runs callback function and returns callback function error.
func (r *Runner) Execute(fn func() error) error {
	if r.attempts > 0 || r.counter == 0 {
		r.counter++
	}
	return fn()
}

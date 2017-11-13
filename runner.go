package repeat

import (
	"context"
	"time"
)

// DefaultAttempts default number of attempts.
const DefaultAttempts = 5

// Reason the execution stop reason.
type Reason string

// Reasons
const (
	AttemptsReached   Reason = "attempts is reached"
	ContextDoneSignal Reason = "received completion signal from the context"
)

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

func (r *Runner) Next() bool {
	if r.counter >= r.attempts {
		r.reason = AttemptsReached
		return false
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

// Execute
func (r *Runner) Execute(fn func() error) error {
	r.counter++
	return fn()
}

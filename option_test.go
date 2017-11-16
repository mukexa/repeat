package repeat

import (
	"context"
	"testing"
	"time"
)

func TestJitter(t *testing.T) {
	cases := []time.Duration{
		0,
		time.Nanosecond,
		time.Second,
		time.Minute,
		time.Hour,
		100000000 * time.Second,
	}

	for _, span := range cases {
		min, max := span*-1, span
		if j := jitter(span); j < min || j > max {
			t.Errorf("jitter(%d) == %d, want in range %d to %d", span, j, min, max)
		}
	}
}

func BenchmarkJitter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		jitter(time.Second)
	}
}

func TestWithBackoffDelay(t *testing.T) {
	cases := []struct {
		startDelay, maxDelay, jitterDelay time.Duration
		wont                              []time.Duration
	}{
		{
			startDelay:  0,
			maxDelay:    time.Hour,
			jitterDelay: 0,
			wont:        []time.Duration{0, 0, 0, 0},
		}, {
			startDelay:  time.Second,
			maxDelay:    6 * time.Second,
			jitterDelay: time.Nanosecond,
			wont:        []time.Duration{time.Second, 2 * time.Second, 4 * time.Second, 6 * time.Second, 6 * time.Second},
		}, {
			startDelay:  9 * time.Minute,
			maxDelay:    time.Hour,
			jitterDelay: time.Second,
			wont:        []time.Duration{9 * time.Minute, 18 * time.Minute, 36 * time.Minute, time.Hour, time.Hour},
		}, {
			startDelay:  2 * time.Nanosecond,
			maxDelay:    6 * time.Nanosecond,
			jitterDelay: time.Second,
			wont:        []time.Duration{2 * time.Nanosecond, 4 * time.Nanosecond, 6 * time.Nanosecond, 6 * time.Nanosecond, 6 * time.Nanosecond},
		},
	}

	for _, c := range cases {
		r := NewRunner(context.Background(),
			WithBackoffDelay(c.startDelay, c.maxDelay, c.jitterDelay))
		for _, wont := range c.wont {
			min, max := wont-c.jitterDelay, wont+c.jitterDelay
			if d := r.delay(); d < min || d > max {
				t.Errorf("WithBackoffDelay got %d, want in range %d to %d", d, min, max)
			}
		}
	}
}

func BenchmarkBackoffDelay(b *testing.B) {
	r := NewRunner(context.Background(),
		WithBackoffDelay(time.Second, time.Hour, time.Second))
	for n := 0; n < b.N; n++ {
		r.delay()
	}
}

func TestWithDelay(t *testing.T) {
	const probeCount = 10

	cases := []time.Duration{
		time.Nanosecond,
		time.Second,
		time.Hour,
	}
	for _, delay := range cases {
		r := NewRunner(context.Background(),
			WithDelay(delay))

		for i := 0; i < probeCount; i++ {
			if d := r.delay(); d != delay {
				t.Errorf("WithDelay got %d, want %d", d, delay)
			}
		}

	}
}

func TestWithDelayFunc(t *testing.T) {
	const probeCount = 10

	cases := []time.Duration{
		time.Nanosecond,
		time.Second,
		time.Hour,
	}
	for _, delay := range cases {
		r := NewRunner(context.Background(),
			WithDelayFunc(func() time.Duration { return delay }))

		for i := 0; i < probeCount; i++ {
			if d := r.delay(); d != delay {
				t.Errorf("WithDelayFunc got %d, want %d", d, delay)
			}
		}

	}
}

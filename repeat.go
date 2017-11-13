package repeat

import (
	"context"
)

// Do starts repeat execution based on the runner configuration.
// If the callback function returns an error it stops execution and returns the error.
func Do(ctx context.Context, callback func() error, opts ...OptFunc) error {
	r := NewRunner(ctx, opts...)
	for r.Next() {
		err := r.Execute(callback)
		if err != nil {
			return &ExecuteError{Count: r.counter, Err: err}
		}
	}
	return nil
}

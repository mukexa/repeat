package repeat

import (
	"context"

	"github.com/fschnko/errpack"
)

// Retry performs a rerun based on the configuration of the runner
// until it succeeds or stops the runner.
// Returns the success flag and error packet, if any.
func Retry(ctx context.Context, callback func() error, opts ...OptFunc) (bool, error) {
	r := NewRunner(ctx, opts...)

	errp := errpack.New("retry")
	for r.Next() {
		err := r.Execute(callback)
		if err == nil {
			return true, errp.ErrorOrNil()
		}
		errp.Add(err)
	}
	errp.Add(&ExecuteError{Count: r.Count(), Reason: r.Reason()})

	return false, errp.ErrorOrNil()
}

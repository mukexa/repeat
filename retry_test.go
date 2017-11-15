package repeat

import (
	"context"
	"fmt"
	"time"
)

func ExampleRetry() {
	callback := func() (err error) {
		_, err = fmt.Println("Retry example")
		return
	}

	ok, err := Retry(context.Background(), callback)
	if err != nil {
		// handle error
	}
	if !ok {
		//
	}
	// Output:
	// Retry example
}

func ExampleRetry_context() {
	callback := func() error {
		time.Sleep(time.Nanosecond * 2)
		return fmt.Errorf("callback error")
	}

	const timeout = 1 * time.Nanosecond

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	_, err := Retry(ctx, callback)
	cancel()
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// retry :callback error
	// retry :execution stopped after 1 attempt, reason: received completion signal from the context
}

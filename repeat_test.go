package repeat

import (
	"context"
	"errors"
	"fmt"
)

func ExampleDo() {
	callback := func() (err error) {
		_, err = fmt.Println("Do example")
		return
	}

	err := Do(context.Background(), callback)
	if err != nil {
		// handle error
	}
	// Output:
	// Do example
	// Do example
	// Do example
	// Do example
	// Do example
}

func ExampleDo_attempts() {
	callback := func() (err error) {
		_, err = fmt.Println("Do example")
		return
	}

	const attempts = 3

	err := Do(context.Background(), callback,
		WithAttempts(attempts))
	if err != nil {
		// handle error
	}
	// Output:
	// Do example
	// Do example
	// Do example
}

func ExampleDo_error() {
	callback := func() (err error) {
		return errors.New("callback error")
	}

	err := Do(context.Background(), callback)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// execution stopped after 1 attempt, with error: callback error
}

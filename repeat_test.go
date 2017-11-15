package repeat

import (
	"context"
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

func ExampleDo_with_attemps() {
	callback := func() (err error) {
		_, err = fmt.Println("Do example")
		return
	}

	const attemps = 3

	err := Do(context.Background(), callback,
		WithAttemps(attemps))
	if err != nil {
		// handle error
	}
	// Output:
	// Do example
	// Do example
	// Do example
}

/*
Package repeat is a small library for cyclic or retries operations.
Regardless of size, this is a very powerful library.
It helps make the code more readable and elegant.

For example, a code like this:
	const (
		delay    = 5 * time.Second
		attempts = 20
	)

	ticker := time.NewTicker(delay)
	cancel := make(chan struct{})
	for i := 0; i < attempts; i++ {
		select {
		case <-ticker.C:
			err := callback()
			if err != nil {
				// handle error
			}
		case <-cancel:
			ticker.Stop()
			break
		}
	}
can be written as follows:
	const (
		delay    = 5 * time.Second
		attempts = 20
	)

	ctx, cancel := context.WithCancel(context.Background())
	err := repeat.Do(ctx, callback,
		repeat.WithAttempts(attempts),
		repeat.WithDelay(delay))
	if err != nil {
		// handle error
	}
*/
package repeat

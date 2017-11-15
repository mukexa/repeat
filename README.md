# repeat
[![Build Status](https://travis-ci.org/fschnko/repeat.svg?branch=master)](https://travis-ci.org/fschnko/repeat) [![Coverage Status](https://coveralls.io/repos/github/fschnko/repeat/badge.svg?branch=master)](https://coveralls.io/github/fschnko/repeat?branch=master)
---
REPEAT is a small library for cyclic or retries operations.
Regardless of size, this is a very powerful library.
It helps make the code more readable and elegant.

For example, a code like this:
```golang
	const (
		delay   = 5 * time.Second
		attemps = 20
	)

	ticker := time.NewTicker(delay)
	cancel := make(chan struct{})
	for i := 0; i < attemps; i++ {
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
```
can be written as follows:
```golang
	const (
		delay   = 5 * time.Second
		attemps = 20
	)

	ctx, cancel := context.WithCancel(context.Background())
	err := repeat.Do(ctx, callback,
		repeat.WithAttemps(attemps),
		repeat.WithDelay(delay))
	if err != nil {
		// handle error
	}
```
---
### Examples

**Simple repeat** *(5 times as default)*
```golang
	err := repeat.Do(context.Background(), callback)
	if err != nil {
		// handle error
	}
```

**Simple repeat**
```golang
	const attemps = 10
	err := repeat.Do(context.Background(), callback,
		repeat.WithAttemps(attemps),
	)
	if err != nil {
		// handle error
	}
```

**Repeat with context timeout**
```golang
const (
		attemps = 1000
		timeout = 10 * time.Second
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	err := repeat.Do(ctx, callback,
		repeat.WithAttemps(attemps),
	)
	cancel()
	if err != nil {
		// handle error
	}
```

**Repeat with delay**
```golang
const (
		attemps = 100
		delay = 10 * time.Second
	)

	err := repeat.Do(context.Background(), callback,
		repeat.WithAttemps(attemps),
		repeat.WithDelay(delay),
	)
	if err != nil {
		// handle error
	}
```

# repeat
[![GoDoc](https://godoc.org/github.com/fschnko/repeat?status.svg)](https://godoc.org/github.com/fschnko/repeat)    [![Build Status](https://travis-ci.org/fschnko/repeat.svg?branch=master)](https://travis-ci.org/fschnko/repeat) [![Coverage Status](https://coveralls.io/repos/github/fschnko/repeat/badge.svg?branch=master)](https://coveralls.io/github/fschnko/repeat?branch=master)    [![Go Report Card](https://goreportcard.com/badge/github.com/fschnko/repeat)](https://goreportcard.com/report/github.com/fschnko/repeat)

---
REPEAT is a small library for cyclic or retries operations.
Regardless of size, this is a very powerful library.
It helps make the code more readable and elegant.

**Usage:**
```shell
go get github.com/fschnko/repeat
```

For example, a code like this:
```golang
	const (
		delay   = 5 * time.Second
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
```
can be written as follows:
```golang
	const (
		delay   = 5 * time.Second
		attempts = 20
	)

	ctx, cancel := context.WithCancel(context.Background())
	err := repeat.Do(ctx, callback,
		repeat.WithAttempts(attempts),
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
	const attempts = 10
	err := repeat.Do(context.Background(), callback,
		repeat.WithAttempts(attempts),
	)
	if err != nil {
		// handle error
	}
```

**Repeat with context timeout**
```golang
const (
		attempts = 1000
		timeout = 10 * time.Second
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	err := repeat.Do(ctx, callback,
		repeat.WithAttempts(attempts),
	)
	cancel()
	if err != nil {
		// handle error
	}
```

**Repeat with delay**
```golang
const (
		attempts = 100
		delay = 10 * time.Second
	)

	err := repeat.Do(context.Background(), callback,
		repeat.WithAttempts(attempts),
		repeat.WithDelay(delay),
	)
	if err != nil {
		// handle error
	}
```

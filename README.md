# repeat
---
[![Build Status](https://travis-ci.org/fschnko/repeat.svg?branch=master)](https://travis-ci.org/fschnko/repeat)
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
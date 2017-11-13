package repeat

import "fmt"

type ExecuteError struct {
	Reason Reason
	Count  int
	Err    error
}

func (e *ExecuteError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("execution stopped after %d attempt, with error: %s", e.Count, e.Err)
	}
	return fmt.Sprintf("execution stopped after %d attempt, reason: %s", e.Count, e.Reason)
}

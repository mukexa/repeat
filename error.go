// Copyright © 2017 Artem Feschenko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repeat

import "fmt"

// ExecuteError is an error that represents the halting information.
type ExecuteError struct {
	Reason Reason
	Count  int
	Err    error
}

func (e *ExecuteError) Error() string {
	msg := fmt.Sprintf("execution stopped after %d attempt", e.Count)

	switch {
	case e.Err != nil:
		return msg + ", with error: " + e.Err.Error()
	case e.Reason != "":
		return msg + ", reason: " + e.Reason.String()
	default:
		return msg
	}
}

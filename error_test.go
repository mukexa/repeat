package repeat

import (
	"errors"
	"testing"
)

func TestExecuteError(t *testing.T) {
	cases := []struct {
		reason Reason
		count  int
		err    error
		wont   string
	}{
		{
			reason: "",
			count:  0,
			err:    nil,
			wont:   "execution stopped after 0 attempt",
		}, {
			reason: "",
			count:  10,
			err:    nil,
			wont:   "execution stopped after 10 attempt",
		}, {
			reason: ContextDoneSignal,
			count:  10,
			err:    nil,
			wont:   "execution stopped after 10 attempt, reason: received completion signal from the context",
		}, {
			reason: "",
			count:  10,
			err:    errors.New("callback error"),
			wont:   "execution stopped after 10 attempt, with error: callback error",
		},
	}

	for _, c := range cases {
		e := &ExecuteError{c.reason, c.count, c.err}
		if got := e.Error(); got != c.wont {
			t.Errorf("&ExecuteError.Error() got %q, want %q", got, c.wont)
		}
	}
}

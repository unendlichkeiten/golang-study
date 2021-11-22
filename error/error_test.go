package error

import "testing"

func TestError(t *testing.T) {
	// test error.Code() and error.Msg()
	t.Logf("error code: %d, error msg: %s", InvalidParams.Code(), InvalidParams.Msg())

	// test NewError
	if err := NewError(10000005, "connect database error"); err != nil {
		t.Logf("error code: %d, error msg: %s", err.Code(), err.Msg())
	}

	// test WithDetail
	if err := NewError(10000006, "process timeout"); err != nil {
		err = err.WithDetail("process node one timeout", "process node two error")
		t.Logf("error code: %d, error detail: %v", err.Code(), err.Details())
	}

	// test StatusCode()
	t.Logf("error code: %d, error msg: %s, status code: %d",
		InvalidParams.Code(), InvalidParams.Msg(), InvalidParams.StatusCode())
}

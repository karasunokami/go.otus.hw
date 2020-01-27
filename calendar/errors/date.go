package errors

import "fmt"

type ErrDateBusy struct {
	err string
}

func (e *ErrDateBusy) Error() string {
	return fmt.Sprintf("Date error: %s", e.err)
}
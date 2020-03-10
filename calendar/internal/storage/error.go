package storage

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	TimeBusyError      Error = "time is busy"
	EventNotFoundError Error = "event not found"
)

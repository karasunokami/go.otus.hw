package dal

import "time"

type Event struct {
	ID            int
	Title         string
	StartDatetime time.Time
	EndDatetime   time.Time
}

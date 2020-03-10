package calendar

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/event"
	"github.com/karasunokami/go.otus.hw/calendar/internal/storage"
	"time"
)

type Calendar interface {
	GetEvent(id storage.EventId) (event.Event, error)

	CreateEvent(startTime time.Time, stopTime time.Time) (storage.EventId, error)

	UpdateEvent(id storage.EventId, event event.Event) error

	DeleteEvent(id storage.EventId) error
}

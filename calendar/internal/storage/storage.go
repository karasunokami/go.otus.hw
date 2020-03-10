package storage

import "github.com/karasunokami/go.otus.hw/calendar/internal/event"

type EventId int

type Storage interface {
	Create(event event.Event) (EventId, error)

	Delete(id EventId) error

	Update(id EventId, event event.Event) error

	Get(id EventId) (event.Event, error)
}

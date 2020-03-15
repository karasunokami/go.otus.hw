package db

import "github.com/karasunokami/go.otus.hw/calendar/internal/dal"

type EventId int

type Storage interface {
	Create(event dal.Event) (EventId, error)

	Delete(id EventId) error

	Update(id EventId, event dal.Event) error

	Get(id EventId) (dal.Event, error)
}

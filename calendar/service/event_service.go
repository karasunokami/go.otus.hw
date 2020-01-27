package service

import "github.com/karasunokami/go.otus.hw/calendar/dal"

type EventService interface {
	// Get one event by id
	Get(id int) (*dal.Event, error)

	// Get list of events by params
	GetList(params map[string]string) ([]*dal.Event, error)

	// Create new event
	Create(e *dal.Event) error

	// Remove event by id
	Remove(id int) error

	// Update event by id
	Update(id int, data map[string]string) error
}

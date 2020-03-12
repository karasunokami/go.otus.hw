package storage

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/event"
	"time"
)

type Client struct {
	events map[EventId]event.Event
}

func NewClient() Client {
	return Client{events: make(map[EventId]event.Event)}
}

func (c *Client) Get(id EventId) (event.Event, error) {
	if !c.eventExists(id) {
		return event.Event{}, EventNotFoundError
	}

	return c.events[id], nil
}

func (c *Client) Create(event event.Event) (EventId, error) {
	if busy := !c.isTimeAvailable(event.StartDatetime); busy {
		return 0, TimeBusyError
	}

	eventId := EventId(len(c.events) + 1)
	c.events[eventId] = event

	return eventId, nil
}

func (c *Client) Update(id EventId, event event.Event) error {
	if !c.eventExists(id) {
		return EventNotFoundError
	}

	if busy := !c.isTimeAvailable(event.StartDatetime); busy {
		return TimeBusyError
	}

	c.events[id] = event
	return nil
}

func (c *Client) Delete(id EventId) error {
	if !c.eventExists(id) {
		return EventNotFoundError
	}

	delete(c.events, id)

	return nil
}

func (c Client) isTimeAvailable(time time.Time) bool {
	for _, evt := range c.events {
		if evt.StartDatetime == time {
			return false
		}
	}

	return true
}

func (c Client) eventExists(id EventId) bool {
	_, ok := c.events[id]

	return ok
}

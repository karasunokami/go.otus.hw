package storage

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/event"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var timeFrom = time.Now()

var events = []event.Event{
	{
		Title:         "1",
		StartDatetime: timeFrom,
		EndDatetime:   timeFrom.Add(time.Minute * 30),
	},
	{

		Title:         "2",
		StartDatetime: timeFrom.Add(time.Minute * 30 * 2),
		EndDatetime:   timeFrom.Add(time.Minute*30*2 + time.Minute*30),
	},
}

func TestClientCreate(t *testing.T) {
	c := New()

	_, err := c.Create(events[0])

	assert.Nil(t, err)
}

func TestClientDelete(t *testing.T) {
	c := New()

	id, _ := c.Create(events[0])

	err := c.Delete(id)

	assert.Nil(t, err)
}

func TestClientUpdate(t *testing.T) {
	c := New()

	id, _ := c.Create(events[0])

	err := c.Update(id, events[1])
	assert.Nil(t, err)

	evt, err := c.Get(id)
	assert.Nil(t, err)

	assert.Equal(t, events[1].Title, evt.Title)

}

func TestClientGet(t *testing.T) {
	c := New()

	id, _ := c.Create(events[0])

	evt, err := c.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, events[0].Title, evt.Title)
}

func TestCreateWithSameDate(t *testing.T) {
	c := New()

	_, err := c.Create(events[0])
	assert.Nil(t, err)

	_, err = c.Create(events[0])
	assert.EqualError(t, err, "time is busy")
}

func TestGetNotFound(t *testing.T) {
	c := New()

	_, err := c.Get(1)
	assert.EqualError(t, err, "event not found")
}

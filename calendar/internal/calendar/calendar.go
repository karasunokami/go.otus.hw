package calendar

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/event"
	"github.com/karasunokami/go.otus.hw/calendar/internal/storage"
	"log"
	"time"
)

type Service struct {
	storage storage.Client
}

func NewCalendar() *Service {
	return &Service{storage: storage.NewClient()}
}

func (s *Service) CreateEvent(startDatetime, endDatetime time.Time) (storage.EventId, error) {
	evt := event.Event{
		Title:         "",
		StartDatetime: startDatetime,
		EndDatetime:   endDatetime,
	}

	id, err := s.storage.Create(evt)
	if err != nil {
		log.Printf("Failed to create event: %s", err)
	}

	return id, nil
}

func (s *Service) GetEvent(id storage.EventId) (event.Event, error) {
	evt, err := s.storage.Get(id)
	if err != nil {
		log.Printf("Failed to get event: %s", err)
	}

	return evt, nil
}

func (s *Service) UpdateEvent(id storage.EventId, event event.Event) error {
	return s.storage.Update(id, event)
}

func (s *Service) DeleteEvent(id storage.EventId) error {
	return s.storage.Delete(id)
}

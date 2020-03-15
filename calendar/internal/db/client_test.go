package db

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/dal"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ClientTestSuite struct {
	suite.Suite

	events []dal.Event
	client Client
}

func (s *ClientTestSuite) SetupTest() {

	timeFrom := time.Now()
	s.events = []dal.Event{
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

	s.client = NewClient()
}

func (s *ClientTestSuite) TestClientCreate() {
	_, err := s.client.Create(s.events[0])

	s.NoError(err)
}

func (s *ClientTestSuite) TestClientDelete() {
	id, _ := s.client.Create(s.events[0])
	err := s.client.Delete(id)

	s.NoError(err)
}

func (s *ClientTestSuite) TestClientUpdate() {
	id, _ := s.client.Create(s.events[0])

	err := s.client.Update(id, s.events[1])
	s.NoError(err)

	evt, err := s.client.Get(id)
	s.NoError(err)

	s.Equal(s.events[1].Title, evt.Title)
}

func (s *ClientTestSuite) TestClientGet() {
	id, _ := s.client.Create(s.events[0])

	evt, err := s.client.Get(id)
	s.NoError(err)

	s.Equal(s.events[0].Title, evt.Title)
}

func (s *ClientTestSuite) TestCreateWithSameDate() {
	_, err := s.client.Create(s.events[0])
	s.NoError(err)

	_, err = s.client.Create(s.events[0])
	s.EqualError(err, "time is busy")
}

func (s *ClientTestSuite) TestGetNotFound() {
	_, err := s.client.Get(1)
	s.EqualError(err, "dal not found")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

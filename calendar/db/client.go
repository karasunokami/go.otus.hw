package db

import "github.com/karasunokami/go.otus.hw/calendar/config"

type Client interface {
	Create(config.DB) *Client
}
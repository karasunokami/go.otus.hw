package app

import (
	"log"
	"sync"
)

type Container struct {
	mux      sync.RWMutex
	services map[string]interface{}
}

type Provider interface {
	Register(c *Container)
}

func NewContainer() *Container {
	return &Container{
		mux:      *new(sync.RWMutex),
		services: make(map[string]interface{}),
	}
}

func (c *Container) Set(id string, method func(c *Container) interface{}) {
	if c.services == nil {
		c.services = make(map[string]interface{})
	}

	c.mux.RLock()
	c.services[id] = method(c)
	c.mux.RUnlock()
}

func (c *Container) Register(provider Provider) {
	provider.Register(c)
}

func (c *Container) Remove(id string) Container {
	c.mux.RLock()

	delete(c.services, id)

	c.mux.RUnlock()

	return *c
}

func (c *Container) Get(id string) interface{} {
	c.mux.RLock()

	_, ok := c.services[id]

	c.mux.RUnlock()

	if !ok {
		log.Panicf("todo to log: service %s undefined", id)
	}

	return c.services[id]
}

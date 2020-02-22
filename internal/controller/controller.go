package controller

import (
	"go-husky/internal/log"
	"go-husky/internal/server"
)

var logger = log.GetLogger()

type Controller struct {
	requestMappings []server.RequestMapping
}

func (c *Controller) registerRequestMapping(mappings ...server.RequestMapping)  {
	if c.requestMappings == nil {
		c.requestMappings = []server.RequestMapping{}
	}
	c.requestMappings = append(c.requestMappings, mappings...)
}

func (c *Controller) GetRequestMappings() []server.RequestMapping {
	return c.requestMappings
}
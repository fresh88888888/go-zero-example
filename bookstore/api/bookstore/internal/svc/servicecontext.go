package svc

import (
	"umbrella.github.com/go-zero-example/bookstore/api/bookstore/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

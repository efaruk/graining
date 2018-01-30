package main

import (
	"github.com/efaruk/graining/poc/core"
	"github.com/efaruk/graining/poc/modules"
)

// DependencyContainer is dependency container
type DependencyContainer struct {
	Logger         core.Logger
	HostFileModule modules.HostFileModule
}

// NewDependencyContainer as Constructor
func NewDependencyContainer() *DependencyContainer {
	c := &DependencyContainer{}
	c.Logger = core.NewConsoleLogger()
	c.HostFileModule = modules.NewHostFileModule()
	return c
}

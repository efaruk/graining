package modules

import (
	"github.com/efaruk/graining/core"
)

// HostFile ...
type HostFile struct {
	data string
}

// Run ...
func (h HostFile) Run() int {
	return 0
}

// SetContext ...
func (h HostFile) SetContext(c core.Context) int {
	return 0
}

// GetResults ..
func (h HostFile) GetResults() string {
	return "localhost 127.0.0.1"
}

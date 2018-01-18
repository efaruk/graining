package core

// ContextProvider for shared abstractions
type ContextProvider interface {
}

// Context shared memory
type Context struct {
}

// Module is core module interface
type Module interface {
	Run() int
	SetContext(c Context) int
	GetResults() string
}

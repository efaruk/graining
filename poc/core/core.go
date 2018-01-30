package core

import (
	"io"
	"log"
	"os"
)

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
	// Should return Json
	GetResults() string
}

// Logger is core logging interface
type Logger interface {
	// Equivalent to
	Fatalln(v ...interface{})
	Panicln(v ...interface{})
	SetOutput(w io.Writer)
	Println(v ...interface{})
}

// ConsoleLogger default console logger
type consoleLogger struct {
}

// NewConsoleLogger as constructor
func NewConsoleLogger() Logger {
	ncl := new(consoleLogger)
	ncl.SetOutput(os.Stdout)
	return ncl
}

// Fatalln log.Fatalln
func (cl consoleLogger) Fatalln(v ...interface{}) {
	log.Fatalln(v)
}

// Panicln log.Panicln
func (cl consoleLogger) Panicln(v ...interface{}) {
	log.Panicln(v)
}

// SetOutput log.SetOutput
func (cl consoleLogger) SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// Println log.Println
func (cl consoleLogger) Println(v ...interface{}) {
	log.Println(v)
}

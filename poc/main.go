package main

import (
	"fmt"
	"time"

	"github.com/marcsauter/single"
)

func main() {
	// Assure only single instance of this application is running
	s := single.New("autoc")
	s.Lock()
	defer s.Unlock()

	// Dependency Container
	dc := NewDependencyContainer()
	fmt.Println("Hello World!")
	dc.HostFileModule.Run()
	fmt.Println("Host File: ", dc.HostFileModule.GetResults())

	fmt.Println("Exiting: ", time.Now().UTC().Format(time.RFC3339))

}

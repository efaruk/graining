package main

import (
	"fmt"
	"time"

	"github.com/facebookgo/inject"
	"github.com/marcsauter/single"

	"github.com/efaruk/graining/core"
	"github.com/efaruk/graining/modules"
)

type autocApp struct {
	HostFileModule *modules.HostFileModule `inject:"modules"`
}

func main() {
	// Assure only single instance of this application is running
	s := single.New("autoc")
	s.Lock()
	defer s.Unlock()

	// DI Container Graph
	var g inject.Graph
	fmt.Println(g)
	var hm core.Module = modules.HostFile{}
	fmt.Println("Hello World!")
	hm.Run()
	fmt.Println("Host File: ", hm.GetResults())

	fmt.Println("Exiting: ", time.Now().UTC().Format(time.RFC3339))

}

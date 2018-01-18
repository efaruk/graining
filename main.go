package main

import (
	"fmt"

	"github.com/efaruk/graining/core"
	"github.com/efaruk/graining/modules"
)

func main() {

	var hm core.Module = modules.HostFile{}
	fmt.Println("Hello World!")
	hm.Run()
	fmt.Println("Host File: ", hm.GetResults())
}

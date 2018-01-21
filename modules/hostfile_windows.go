// +build windows

package modules

import (
	"os"
)

var hostsFilePath = os.Getenv("SystemRoot") + `\System32\drivers\etc\hosts`

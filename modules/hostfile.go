package modules

import (
	"io/ioutil"
	"strings"

	"github.com/efaruk/graining/core"
)

type HostFileModule interface {
	core.Module
	ParseHostFile(hostFileContent string)
}

// HostFile ...
type HostFile struct {
	data  string
	ipmap map[string][]string
}

type HostFileEntryList struct {
}

// Run core.Module Run implementiation
func (h HostFile) Run() int {
	bs, _ := readHostsFile(hostsFilePath)
	parseHosts(bs)
	return 0
}

// SetContext core.Module SetContext implementiation
func (h HostFile) SetContext(c core.Context) int {
	return 0
}

// GetResults core.Module GetResults implementation
func (h HostFile) GetResults() string {
	return "localhost 127.0.0.1"
}

func readHostsFile(path string) ([]byte, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func parseHosts(hostsFileContent []byte) map[string][]string {
	if hostsFileContent == nil {
		return nil
	}
	hostsMap := map[string][]string{}
	for _, line := range strings.Split(strings.Trim(string(hostsFileContent), " \t\r\n"), "\n") {
		line = strings.Replace(strings.Trim(line, " \t"), "\t", " ", -1)
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}
		pieces := strings.SplitN(line, " ", 2)
		if len(pieces) > 1 && len(pieces[0]) > 0 {
			if names := strings.Fields(pieces[1]); len(names) > 0 {
				if _, ok := hostsMap[pieces[0]]; ok {
					hostsMap[pieces[0]] = append(hostsMap[pieces[0]], names...)
				} else {
					hostsMap[pieces[0]] = names
				}
			}
		}
	}
	return hostsMap
}

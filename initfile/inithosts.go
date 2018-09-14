package initfile

import (
	"github.com/karamelisthecat/karamel/flag"
	"github.com/karamelisthecat/karamel/hostsfile"
	"os"
)

var pathHosts = "/etc/hosts"

func InitHosts() {
	hostsfile.LinesHost, _ = hostsfile.ReadHostFile(pathHosts)
	hostsfile.FindGroupNames()
	if len(os.Args) > 1 {
		flag.OneFlag()
	} else {
		userInterface()
	}
}

func userInterface() {
	var isRunning bool
	isRunning = true
	for isRunning {
		isRunning = hostsfile.UserOptMenu()
	}
}

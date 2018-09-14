package initfile

import (
	"github.com/karamelisthecat/karamel/hostsfile"
)

var PathHosts = "/etc/hosts"

func InitHosts() {
	hostsfile.LinesHost, _ = hostsfile.ReadHostFile(PathHosts)
	hostsfile.FindGroupNames()
	userInterface()
}

func userInterface() {
	var isRunning bool
	isRunning = true
	for isRunning {
		isRunning = hostsfile.UserOptMenu()
	}
}

package initfile

import (
	//"fmt"
	"github.com/karamelisthecat/karamel/hostsFile/flag"
	"github.com/karamelisthecat/karamel/hostsFile/hostsfile"
	"os"
)

func InitHosts() {
	hostsfile.LinesHost, _ = hostsfile.ReadHostFile("/etc/hosts")
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

package initfile

import (
	"fmt"
	"github.com/karamelisthecat/karamel/hostsFile/flag"
	"github.com/karamelisthecat/karamel/hostsFile/hostsfile"
)

func InitHosts() {
	hostsfile.LinesHost, _ = hostsfile.ReadHostFile("/etc/hosts")
	hostsfile.FindGroupNames()
	flag.OneFlag()
	userInterface()
}

func userInterface() {
	var isRunning bool
	isRunning = true
	fmt.Println("\n/etc/hosts file")
	fmt.Println("----------------")
	for isRunning {
		isRunning = hostsfile.UserOptMenu()
	}
}

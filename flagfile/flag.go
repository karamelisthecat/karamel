package flagfile

import (
	"flag"
	"github.com/karamelisthecat/karamel/hostsfile"
	"github.com/karamelisthecat/karamel/initfile"
	resolv "github.com/karamelisthecat/karamel/resolvconf"
)

var (
	Addns *string
	Delns *int
)

func Ifflags() {
	addIPptr := flag.Bool("addIP", false, "To add IP address.")
	addGroupPtr := flag.Bool("addGroup", false, "To add group.")
	addWebPtr := flag.Bool("web", false, "To web interface.")
	//	for {
	Addns = flag.String("addnameserver", "0.0.0.0", "for add nameserver")
	Delns = flag.Int("delnameserver", 0, "for delete nameserver")
	flag.Parse()
	//    flag.Parse()
	if *addIPptr {
		hostsfile.AddIPblock()
	}
	if *addGroupPtr {
		hostsfile.AddGroup()
	}
	if *addWebPtr {
		initfile.WebInterface()
	}

	if *Addns != "0.0.0.0" && *Delns != 0 {
		resolv.AddNameserver(Addns)
		resolv.DeleteNameserver(Delns)
		//			break
	} else if *Delns != 0 {
		resolv.DeleteNameserver(Delns)
		//			break
	} else if *Addns != "0.0.0.0" {
		resolv.AddNameserver(Addns)
		//			break
	}
	//	}
}

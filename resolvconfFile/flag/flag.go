package flag

import (
	"flag"
	resolv "github.com/karamelisthecat/karamel/resolvconfFile/resolvconf"
)

var (
	Addns *string
	Delns *int
)

func Ifflags() {
	Addns = flag.String("addnameserver", "0.0.0.0", "for add nameserver")
	Delns = flag.Int("delnameserver", 0, "for delete nameserver")
	flag.Parse()
	if *Addns != "0.0.0.0" && *Delns != 0 {

		resolv.AddNameserver(Addns)
		resolv.DeleteNameserver(Delns)

	} else if *Delns != 0 {
		resolv.DeleteNameserver(Delns)
	} else if *Addns != "0.0.0.0" {
		resolv.AddNameserver(Addns)
	} else {
		resolv.SelectMenu()
	}
}

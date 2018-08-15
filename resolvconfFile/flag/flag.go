package flag

import (
	resolv "../resolvconf"
	"flag"
)

var (
	Addns *string
	Delns *int
)

func Ifflags() {
	Addns = flag.String("addnameserver", "0.0.0.0", "for add nameserver")
	Delns = flag.Int("delnameserver", 0, "for delete nameserver")
	flag.Parse()
	if *Addns != "0.0.0.0" {
		resolv.AddNameserver(Addns)
	}
	if *Delns != 0 {
		resolv.DeleteNameserver(Delns)
	} else {
		resolv.SelectMenu()
	}
}
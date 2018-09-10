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
	for {
		Addns = flag.String("addnameserver", "0.0.0.0", "for add nameserver")
		Delns = flag.Int("delnameserver", 0, "for delete nameserver")
		flag.Parse()
		if *Addns != "0.0.0.0" && *Delns != 0 {
			resolv.AddNameserver(Addns)
			resolv.DeleteNameserver(Delns)
			break
		} else if *Delns != 0 {
			resolv.DeleteNameserver(Delns)
			break
		} else if *Addns != "0.0.0.0" {
			resolv.AddNameserver(Addns)
			break
		}
	}
}

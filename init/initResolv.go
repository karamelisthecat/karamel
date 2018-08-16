package init

import (
	"github.com/karamelisthecat/karamel/resolvconfFile/flag"
	"github.com/karamelisthecat/karamel/resolvconfFile/resolvconf"
)

func initResolv() {
	resolvconf.OpenReadFile()
	resolvconf.KeepResolvconf()
	flag.Ifflags()
	resolvconf.SaveChange()

}

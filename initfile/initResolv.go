package initfile

import (
	"github.com/karamelisthecat/karamel/resolvconfFile/flag"
	"github.com/karamelisthecat/karamel/resolvconfFile/resolvconf"
)

func InitResolv() {
	resolvconf.OpenReadFile()
	resolvconf.KeepResolvconf()
	flag.Ifflags()
	resolvconf.SaveChange()

}

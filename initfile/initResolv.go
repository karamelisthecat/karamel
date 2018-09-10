package initfile

import (
	"github.com/karamelisthecat/karamel/resolvconfFile/flag"
	"github.com/karamelisthecat/karamel/resolvconfFile/resolvconf"
	"os"
)

func InitResolv() {
	resolvconf.OpenReadFile()
	resolvconf.KeepResolvconf()
	if len(os.Args) > 1 {
		flag.Ifflags()
	} else {
		resolvconf.SelectMenu()
	}
	resolvconf.SaveChange()

}

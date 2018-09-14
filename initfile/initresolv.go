package initfile

import (
	"github.com/karamelisthecat/karamel/flag"
	"github.com/karamelisthecat/karamel/resolvconf"
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

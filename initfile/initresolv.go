package initfile

import (
	"github.com/karamelisthecat/karamel/resolvconf"
)

func InitResolv() {
	resolvconf.OpenReadFile()
	resolvconf.KeepResolvconf()
	resolvconf.SelectMenu()
	resolvconf.SaveChange()

}

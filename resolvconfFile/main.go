package main

import (
	flag "./flag"
	resolv "./resolvconf"
)

func main() {
	resolv.OpenReadFile()
	resolv.KeepResolvconf()
	flag.Ifflags()
	resolv.SaveChange()

}

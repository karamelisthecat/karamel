package flag

import (
	"flag"
	"github.com/berfinsari/projeHost/hostsfile"
)

func OneFlag() {
	addIPptr := flag.Bool("addIP", false, "IP eklemek için kullanılır.")
	addGroupPtr := flag.Bool("addGroup", false, "Grup eklemek için kullanılır.")
	flag.Parse()
	if *addIPptr {
		hostsfile.AddIPblock()
	}
	if *addGroupPtr {
		hostsfile.AddGroup()
	}
}
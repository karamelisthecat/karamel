package main

import (
	"fmt"
	"github.com/karamelisthecat/karamel/initfile"
)

func main() {
Loop:
	for {
		var choice string
		fmt.Println("which file would you like to change: ")
		fmt.Println("1: hosts file")
		fmt.Println("2: resolvconf file")
		fmt.Println("q: exit")
		entry, _ := fmt.Scanf("%s", &choice)
		if entry == 0 {
			continue
		}
		switch choice {
		case "1":
			initfile.InitHosts()
		case "2":
			initfile.InitResolv()
		case "q":
			break Loop
		default:
			fmt.Println("please enter valid value")
		}
	}
}

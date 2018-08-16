package main

import (
	"fmt"
	"github.com/karamelisthecat/karamel/initfile"
)

func main() {
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
		if choice == "1" {
			initfile.InitHosts()

		} else if choice == "2" {
			initfile.InitResolv()

		} else if choice == "q" {
			break
		} else {
			fmt.Println("please enter valid value")
		}

	}
}

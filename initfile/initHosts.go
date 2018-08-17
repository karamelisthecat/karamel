package initfile

import (
	"fmt"
	"github.com/karamelisthecat/karamel/hostsFile/flag"
	"github.com/karamelisthecat/karamel/hostsFile/hostsfile"
)

func InitHosts() {
	hostsfile.LinesHost, _ = hostsfile.ReadHostFile("/etc/hosts")
	hostsfile.FindGroupNames()
	flag.OneFlag()
	userInterface()
}

func userInterface() {
	var isRunning bool
	var optionNumber string
	isRunning = true
	fmt.Println("\n/etc/hosts file")
	fmt.Println("----------------")
	for isRunning {
		fmt.Println("______________________________________")
		fmt.Println("1. Add Group.\n2. View Group and IP Addresses.\n3. List Group Names.\n4. Add IP Address.\n5. View Hosts File.\n6. Add IP Field to Group.\n7. Add Alias to IP Address.\n8. Remove Commend Line Tag to Enable IP Address.\nq: Back to Main Menu")
		fmt.Println("______________________________________")
		fmt.Print("\nSeçeneklerden birini seçiniz: ")
		fmt.Scan(&optionNumber)
		fmt.Printf("\n")
		switch optionNumber {
		case "q":
			isRunning = false
			break
		case "1":
			hostsfile.AddGroup()
		case "2":
			hostsfile.ListGroup()
		case "3":
			hostsfile.WriteGroupNames()
		case "4":
			hostsfile.AddIPblock()
		case "5":
			hostsfile.WriteLines()
		case "6":
			hostsfile.AddFieldstoGroup()
		case "7":
			hostsfile.AddAlias()
		case "8":
			hostsfile.RemoveCommendLineIP()
		default:
			fmt.Println("Geçersiz bir işlem girdiniz")
		}
	}
}

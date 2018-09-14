package hostsfile

import "fmt"

func UserOptMenu() bool {
	var optionNumber string
	var isRunning bool
	isRunning = true
	AddUnderscoreLine()
	fmt.Println("»»»»»      /etc/hosts file       «««««")
	AddUnderscoreLine()
	fmt.Println("1. Add Group.\n2. View Group and IP Addresses.\n3. List Group Names.\n4. Add IP Address.\n5. View Hosts File.\n6. Add IP Field to Group.\n7. Add Alias to IP Address.\n8. Remove Commend Line Tag to Enable IP Address.\nq: Back to Main Menu")
	AddUnderscoreLine()
	fmt.Print("Please select one: ")
	fmt.Scan(&optionNumber)
	AddUnderscoreLine()
	fmt.Printf("\n")
	switch optionNumber {
	case "q":
		isRunning = false
		break
	case "1":
		AddGroup()
	case "2":
		ListGroup()
	case "3":
		WriteGroupNames()
		WaitUser()
	case "4":
		AddIPblock()
	case "5":
		WriteLines()
	case "6":
		AddFieldstoGroup()
	case "7":
		AddAlias()
	case "8":
		RemoveCommendLineIP()
	default:
		fmt.Println("You have entered an invalid option.")
		WaitUser()
	}
	return isRunning
}

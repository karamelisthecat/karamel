package hostsfile

import (
	"fmt"
	"net"
	"strings"
)

func UserOptMenu() bool {
	var optionNumber string
	var isRunning bool
	isRunning = true
	fmt.Println("______________________________________")
	fmt.Println("»»»»»      /etc/hosts file       «««««")
	fmt.Println("______________________________________")
	fmt.Println("1. Add Group.\n2. View Group and IP Addresses.\n3. List Group Names.\n4. Add IP Address.\n5. View Hosts File.\n6. Add IP Field to Group.\n7. Add Alias to IP Address.\n8. Remove Commend Line Tag to Enable IP Address.\nq: Back to Main Menu")
	fmt.Println("______________________________________")
	fmt.Print("Please select one: ")
	fmt.Scan(&optionNumber)
	fmt.Println("______________________________________")
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

// Adding ip block to /etc/hosts file.
func AddIPblock() {
	var addingField string
	addingField = ReturnField()
	for i := 0; i < len(LinesHost); i++ {
		if LinesHost[i] == "\n" {
			AddLinesHosts(addingField, i, i)
			fmt.Println("Success!")
			LastViewoftheFile()
			break
		}
	}
}

func ReturnField() string {
	var ipField string
	ipaddress, err := checkIP()
	if err != nil {
		fmt.Print("Error") // **
	}
	hostname := addHostname()
	ipField = ipaddress + "\t" + hostname + "\n"
	return ipField
}

func addHostname() string {
	var hostname string
	var control bool
	control = true
	for control {
		fmt.Print("Please enter the hostname:")
		fmt.Scan(&hostname)
		control, hostname = hostnameCheck(hostname)
	}
	return hostname
}

// Checks if the hostname is used or not used.
func hostnameCheck(hostname string) (bool, string) {
	for i := 0; i < len(LinesHost); i++ {
		if strings.Contains(LinesHost[i], hostname) && strings.Contains(LinesHost[i], "#") != true {
			fmt.Println(hostname, " , already used. Please enter a new one.")
			return true, hostname
		}
	}
	return false, hostname
}

func checkIP() (string, error) {
	var ipv4Addr net.IP
	for {
		var ipv4 string
		fmt.Printf("Please enter the IP address: ")
		fmt.Scanln(&ipv4)
		ipv4Addr = net.ParseIP(ipv4)
		if ipv4Addr.To4() == nil {
			fmt.Println("Invalid IP address.\n") // **
		} else {
			break
		}
	}
	ipv4Str := ipv4Addr.String()
	return ipv4Str, nil
}

func AddGroup() {
	var newgroup string
	var group string
	var control = true
	fmt.Print("Please enter the group name: ")
	fmt.Scan(&group)
	control = groupCheck(group)
	if control {
		newgroup = "# *" + group + "*" + "\n"
		LinesHost = append(LinesHost, newgroup)
		GroupName = append(GroupName, group)
		WriteHostFile(LinesHost)
		AddNewLine()
		setGroup(group)
		for {
			if askOpt() {
				setGroup(group)
			} else {
				break
			}
		}
		fmt.Println("______________________________________")
		fmt.Printf("\nSuccess!\n" + group + " group available!")
		FindtheGroup(group)
		LastViewoftheFile()
	}
}

// Checking group name is available
func groupCheck(group string) bool {
	for i := 0; i < len(GroupName); i++ {
		if GroupName[i] == group {
			fmt.Print("This group exist.") // **
			return false
		}
	}
	if strings.Contains(group, "*") {
		fmt.Print("You can not create a group with '*' charachter in it.")
		return false
	}
	return true
}

// Asking the user.
func askOpt() bool {
	var userOpt string
	fmt.Print("\nWould you like to add more? \n('y' or 'Y'): ")
	fmt.Scan(&userOpt)
	fmt.Print("\n")
	if userOpt == "y" || userOpt == "Y" {
		return true
	}
	return false
}

// Print the last view of the file.
func LastViewoftheFile() {
	var userOpt string
	fmt.Print("\nWould you like to see the last view of the file? \n('y' or 'Y': ")
	fmt.Scan(&userOpt)
	if userOpt == "y" || userOpt == "Y" {
		WriteLines()
	}
}

// Get the group name from the user.
func AddFieldstoGroup() {
	var addGroup string
	WriteGroupNames()
	fmt.Print("\nPlease enter the group name: ")
	fmt.Scan(&addGroup)
	setGroup(addGroup)
	fmt.Println("Success!")
	LastViewoftheFile()
}

// Find the group and add ip field.
func setGroup(addGroup string) {
	var ctrl int
	var emptyLineTemp bool
	var c int
	ctrl = 0
	for i := 0; i < len(LinesHost); i++ {
		check, j := groupControl(i)
		if check && addGroup == GroupName[j] {
			emptyLineTemp, c = FindEmptyLine(i)
			if emptyLineTemp {
				fieldTemp := ReturnField()
				AddLinesHosts(fieldTemp, c, c) // This line writing the file.
				ctrl = 1
				break
			}
		}
	}
	if ctrl == 0 {
		fmt.Println("This group does not exist.")
		var userOpt string
		fmt.Println("Would you want to add this group? \n ('y' or 'Y': ")
		fmt.Scan(&userOpt)
		if userOpt == "y" || userOpt == "Y" {
			AddGroup()
		}
		WaitUser()
	}
}

func groupControl(i int) (bool, int) {
	for j := 0; j < len(GroupName); j++ {
		if strings.HasPrefix(LinesHost[i], "# *") && strings.Contains(LinesHost[i], "*"+GroupName[j]+"*") {
			return true, j
		}
	}
	return false, -1
}

//Adding Alias to specific IP address.
func AddAlias() {
	var addIpTemp string
	LastViewoftheFile()
	ipField, i := checkAlias()
	if i != -1 {
		AddLinesHosts(ipField, i, (i + 1))
		fmt.Println("Success!")
		LastViewoftheFile()
	} else {
		fmt.Printf("Would you like to add this IP address? \n ('y' or 'Y': ")
		fmt.Scan(&addIpTemp)
		AddIPblock()
		LastViewoftheFile()
	}

}

func checkAlias() (string, int) {
	var iptemp string
	var ctrl = false
	var ipField string
	var i int
	fmt.Printf("Which IP address would you like to add alias: ")
	fmt.Scan(&iptemp)
	for i = 0; i < len(LinesHost); i++ {
		if strings.HasPrefix(LinesHost[i], iptemp) {
			fmt.Printf("Please enter the alias: ")
			var addTemp string
			fmt.Scan(&addTemp)
			bolTemp := strings.Split(string(LinesHost[i]), "\n")
			ipField = string(bolTemp[0]) + "\t" + addTemp + "\n"
			ctrl = true
			break
		}
	}
	if ctrl != true {
		fmt.Println("This IP address does not exist.")
		return iptemp, -1
	}
	return ipField, i
}

// Waiting user.
func WaitUser() {
	var temp string
	entry, _ := fmt.Scanf("%s", &temp)
	if entry != 0 {
		fmt.Println("\n")
	}
}

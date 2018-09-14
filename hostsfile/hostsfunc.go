package hostsfile

import (
	"fmt"
	"net"
	"strings"
)

// Adding ip block to /etc/hosts file.
func AddIPblock() {
	var ipaddress string
	var hostname string
	ipaddress, hostname = AskIPaddressToUser()
	ctrl := ReturnipField(ipaddress, hostname)
	if ctrl {
		WriteHostFile(LinesHost)
		fmt.Println("\nSuccess !")
		LastViewoftheFile()
	}
}

// Asking ip address and hostname to user.
func AskIPaddressToUser() (string, string) {
	var ctrl = true
	var control = true
	var ipv4, hostname string
	for ctrl || control {
		for ctrl {
			fmt.Printf("Please enter the IP address: ")
			fmt.Scanln(&ipv4)
			ipv4, ctrl = ctrlIP(ipv4)
		}
		for control {
			fmt.Print("Please enter the hostname:")
			fmt.Scan(&hostname)
			hostname, control = ctrlHostname(hostname)
		}
	}
	return ipv4, hostname
}

func ReturnipField(ipaddress string, hostname string) bool {
	var ipField string
	ipField = ipaddress + "\t" + hostname + "\n"
	for i := 0; i < len(LinesHost); i++ {
		if LinesHost[i] == "\n" {
			AddLinesHosts(ipField, i, i)
			return true
		}
	}
	return false
}

// Check if the IP address is valid or invalid.
func ctrlIP(ipv4 string) (string, bool) {
	var ipv4Addr net.IP
	ipv4Addr = net.ParseIP(ipv4)
	if ipv4Addr.To4() == nil {
		fmt.Println("Invalid IP address.\n")
		return "", true
	}
	ipv4Str := ipv4Addr.String()
	return ipv4Str, false
}

// Checks if the hostname is used or not used.
func ctrlHostname(hostname string) (string, bool) {
	for i := 0; i < len(LinesHost); i++ {
		if strings.Contains(LinesHost[i], hostname) && strings.Contains(LinesHost[i], "#") != true {
			fmt.Println(hostname, " , already used. Please enter a new one.")
			return hostname, true
		}
	}
	return hostname, false
}

func AddGroup() {
	var groupname string
	groupname = askGroupnameToUser()
	WriteGroupnameToFile(groupname)
	setGroup(groupname)
	for {
		if askOpt() {
			setGroup(groupname)
		} else {
			break
		}
	}
	fmt.Println("Success!")
	LastViewoftheFile()
}

func AddGroupInterface(gName string, iAddress string, hName string) {
	var ipField string
	WriteGroupnameToFile(gName)
	var c = findGroup(gName)
	if c == -1 {
		fmt.Println("ERROR")
	}
	ipField = iAddress + "\t" + hName + "\n"
	AddLinesHosts(ipField, c, c)
}

func findGroup(groupname string) int {
	var emptyLineTemp bool
	var c int
	for i := 0; i < len(LinesHost); i++ {
		check, j := findGroupLine(i)
		if check && groupname == GroupName[j] {
			emptyLineTemp, c = FindEmptyLine(i)
			if emptyLineTemp {
				return c
			}
		}
	}
	return -1
}

func WriteGroupnameToFile(groupname string) {
	var newgroup string
	newgroup = "# *" + groupname + "*" + "\n"
	LinesHost = append(LinesHost, newgroup)
	GroupName = append(GroupName, groupname)
	WriteHostFile(LinesHost)
	AddNewLine()
}

// Checking group name is available
func ctrlGroupname(groupname string) bool {
	for i := 0; i < len(GroupName); i++ {
		if GroupName[i] == groupname {
			fmt.Print("This group exist.") // **
			return false
		}
	}
	if strings.Contains(groupname, "*") {
		fmt.Print("You can not create a group with '*' charachter in it.")
		return false
	}
	return true
}

// Ask user to group name.
func askGroupnameToUser() string {
	var groupname string
	var ctrl bool
	ctrl = false
	for ctrl != true {
		fmt.Print("Please enter the group name: ")
		fmt.Scan(&groupname)
		ctrl = ctrlGroupname(groupname)
	}
	return groupname
}

// Asking the user for adding more ip fields.
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
func setGroup(groupname string) {
	var ctrl bool
	var emptyLineTemp bool
	var c int
	ctrl = false
	for i := 0; i < len(LinesHost); i++ {
		check, j := findGroupLine(i)
		if check && groupname == GroupName[j] {
			emptyLineTemp, c = FindEmptyLine(i)
			if emptyLineTemp {
				fieldTemp := ReturnField()
				AddLinesHosts(fieldTemp, c, c) // This line writing the file.
				ctrl = true
				break
			}
		}
	}
	if ctrl == false {
		askforAddGroup()
	}
}

func askforAddGroup() {
	fmt.Println("This group does not exist.")
	var userOpt string
	fmt.Println("Would you want to add this group? \n ('y' or 'Y': ")
	fmt.Scan(&userOpt)
	if userOpt == "y" || userOpt == "Y" {
		AddGroup()
	}
	WaitUser()
}

// Ask ip field to user for group.
func ReturnField() string {
	var ipField string
	ipaddress, hostname := AskIPaddressToUser()
	ipField = ipaddress + "\t" + hostname + "\n"
	return ipField
}

func findGroupLine(i int) (bool, int) {
	for j := 0; j < len(GroupName); j++ {
		if strings.HasPrefix(LinesHost[i], "# *") && strings.Contains(LinesHost[i], "*"+GroupName[j]+"*") {
			return true, j
		}
	}
	return false, -1
}

// Adding Alias to specific IP address.
func AddAlias() {
	var i int
	var temp string
	LastViewoftheFile()
	temp = askAliasIPtoUser("askip")
	i = findIPAddressinfile(temp)
	if i != -1 {
		temp = askAliasIPtoUser("askalias")
		ipfieldtemp := AddfiletoAlias(temp, i)
		fmt.Println("Success!")
		fmt.Println(ipfieldtemp)
	} else {
		temp = askAliasIPtoUser("askaddip")
		if temp == "y" || temp == "Y" {
			AddIPblock()
		}
	}
}

// Adding Alias to specific IP address.
func AddAliasInterface(ipaddress string, alias string) {
	var i int
	i = findIPAddressinfile(ipaddress)
	if i == -1 {
		fmt.Println("ERROR")
	}
	AddfiletoAlias(alias, i)
}

func findIPAddressinfile(iptemp string) int {
	for i := 0; i < len(LinesHost); i++ {
		if strings.HasPrefix(LinesHost[i], iptemp) {
			return i
		}
	}
	return -1
}

func AddfiletoAlias(alias string, i int) string {
	var ipField string
	splitTemp := strings.Split(string(LinesHost[i]), "\n")
	ipField = string(splitTemp[0]) + "\t" + alias + "\n"
	AddLinesHosts(ipField, i, (i + 1))
	return ipField
}

func askAliasIPtoUser(opt string) string {
	var addtemp string
	if opt == "askip" {
		fmt.Printf("Which IP address would you like to add alias: ")
		fmt.Scan(&addtemp)
	} else if opt == "askalias" {
		fmt.Printf("Please enter the alias: ")
		fmt.Scan(&addtemp)
	} else {
		fmt.Printf("Would you like to add this IP address? \n ('y' or 'Y'): ")
		fmt.Scan(&addtemp)
	}
	return addtemp
}

// Waiting user.
func WaitUser() {
	var temp string
	entry, _ := fmt.Scanf("%s", &temp)
	if entry != 0 {
		fmt.Println("\n")
	}
}

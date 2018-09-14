package hostsfile

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	filePath  = "/etc/hosts"
	GroupName []string
	LinesHost []string
)

//Reading Hosts file.
func ReadHostFile(filePath string) ([]string, error) {
	var line []string
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print("\nError reading hosts file: ", err)
		return nil, err
	}
	line = strings.SplitAfter(string(file), "\n")
	return line, nil
}

func FindGroupNames() {
	for i := 0; i < len(LinesHost); i++ {
		if strings.HasPrefix(LinesHost[i], "# *") {
			temp := strings.Split(LinesHost[i], "*")
			GroupName = append(GroupName, temp[1])
		}
	}
}

//Writing at the Hosts file.
func WriteHostFile(lines []string) error {
	var alan string
	for i := 0; i < len(lines); i++ {
		alan = alan + lines[i]
	}
	dataBytes := []byte(alan)
	err := ioutil.WriteFile(filePath, dataBytes, 0)
	if err != nil {
		fmt.Print("Error writing hosts file")
		return err
	}
	return nil
}

// Adding a new line at the end of the Hosts file.
func AddNewLine() {
	LinesHost = append(LinesHost, "\n")
	WriteHostFile(LinesHost)
}

// Writing /etc/Hosts file to the screen.
func WriteLines() {
	fmt.Println("\nHosts File ")
	AddUnderscoreLine()
	fmt.Print("\n")
	for i := 0; i < len(LinesHost); i++ {
		fmt.Print(LinesHost[i])
	}
	WaitUser()
}

// Writing /etc/Hosts file to the screen with number.
func writeHostFilewithNmbr() {
	for i := 0; i < len(LinesHost); i++ {
		fmt.Print((i + 1), " ", LinesHost[i])
	}
}

// Writing Group Names to the screen.
func WriteGroupNames() {
	if len(GroupName) == 0 {
		fmt.Println("There is no group in hosts file.")
	} else {
		fmt.Println("»»»»» Groups «««««")
		for i := 0; i < len(GroupName); i++ {
			fmt.Println(GroupName[i])
		}
	}
}

// Writing Group with all fields.
func ListGroup() {
	var nameGroup string
	WriteGroupNames()
	fmt.Print("Which group would you like to wiev: ")
	fmt.Scan(&nameGroup)
	AddUnderscoreLine()
	FindtheGroup(nameGroup)
}

func FindtheGroup(nameGroup string) {
	control := 0
	temp := "# *" + nameGroup + "*"
	for i := 0; i < len(LinesHost); i++ {
		if strings.HasPrefix(LinesHost[i], "# *") && strings.Contains(LinesHost[i], temp) {
			fmt.Print("\n")
			for c := i; c < len(LinesHost); c++ {
				if LinesHost[c] == "\n" {
					control = 1
					break
				}
				fmt.Print(LinesHost[c])
			}
			if control == 1 {
				break
			}
		}
	}
	if control == 0 {
		fmt.Println("This group does not exist.")
	}
	WaitUser()

}

// Find empty line in Hosts file.
func FindEmptyLine(i int) (bool, int) {
	for c := i; c < len(LinesHost); c++ {
		if LinesHost[c] == "\n" {
			return true, c
		}
	}
	return false, -1
}

// Appending new slice to hosts file slice.
func AddLinesHosts(fieldTemp string, before int, after int) {
	var change []string
	change = append(change, LinesHost[:before]...)
	change = append(change, fieldTemp)
	LinesHost = append(change, LinesHost[after:]...)
	WriteHostFile(LinesHost)
}

// Remove command line tag with line number.
func RemoveCommendLineIP() {
	var lineNmbr int
	var newField string
	writeHostFilewithNmbr()
	lineNmbr = checkUserInput()
	newField = deleteCommendLine(lineNmbr - 1)
	AddLinesHosts(newField, (lineNmbr - 1), (lineNmbr))
	AddUnderscoreLine()
	fmt.Println("Success!")
	LastViewoftheFile()
	WaitUser()
}

// Checking the line to provide conditions.
func checkUserInput() int {
	var lines string
	var LineNmbr int
	for {
		var err error
		fmt.Printf("\nPlease enter the line number: ")
		entry, _ := fmt.Scanf("%s", &lines)
		LineNmbr, err = strconv.Atoi(lines)
		if err != nil || entry == 0 || LineNmbr > len(LinesHost) || LineNmbr < 1 {
			fmt.Println("Invalid line number.")
			continue
		} else if strings.HasPrefix(LinesHost[LineNmbr-1], "#") != true || LinesHost[LineNmbr-1] == "" {
			fmt.Println("This line is not a command line.")
			continue
		} else if strings.HasPrefix(LinesHost[LineNmbr-1], "#") && strings.Contains(LinesHost[LineNmbr-1], "\t") != true {
			fmt.Println("This line does not contain an IP field.")
			continue
		}
		break
	}
	return LineNmbr
}

// Print the last view of the file.
func LastViewoftheFile() {
	var userOpt string
	fmt.Print("Would you like to see the last view of the file? \n('y' or 'Y'): ")
	fmt.Scan(&userOpt)
	if userOpt == "y" || userOpt == "Y" {
		WriteLines()
	}
}

func deleteCommendLine(lineNmbr int) string {
	temp := strings.Split(string(LinesHost[lineNmbr]), "#")
	return string(temp[1])
}

func AddUnderscoreLine() {
	fmt.Println("______________________________________")
}

package hostsfile

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	filePath  = "/etc/hosts"
	GroupName []string
	LinesHost []string
)

func ReadHostFile(filePath string) ([]string, error) { //dosyayı okur
	var line []string
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print("\ndosya okuma hatası: ", err)
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
		fmt.Print("ERROR")
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
	fmt.Println("\n Hosts File \n")
	for i := 0; i < len(LinesHost); i++ {
		fmt.Print(LinesHost[i])
	}
	fmt.Print("\n")
}

// Writing Group Names to the screen.
func WriteGroupNames() {
	if len(GroupName) == 0 {
		fmt.Println("Hiç grup bulunmamaktadır.")
	} else {
		fmt.Println("~~ Groups: ~~")
		for i := 0; i < len(GroupName); i++ {
			fmt.Println(GroupName[i])
		}
	}
}

// Writing Group with all fields.
func ListGroup() {
	var nameGroup string
	fmt.Print("Which group would you like to wiev: ")
	fmt.Scan(&nameGroup)
	control := 0
	temp := "# *" + nameGroup + "*"
	for i := 0; i < len(LinesHost); i++ {
		if strings.HasPrefix(LinesHost[i], "# *") && strings.Contains(LinesHost[i], temp) {
			fmt.Println(nameGroup, " grubunun alanları:") //**
			for c := i + 1; c < len(LinesHost); c++ {
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
		fmt.Println("Böyle bir grup bulunmamaktadır")
	}
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

//append işlemlerini yaparak dizileri birleştirir ve dosyaya yazar
func AddLinesHosts(fieldTemp string, before int, after int) {
	var change []string
	change = append(change, LinesHost[:before]...)
	change = append(change, fieldTemp)
	LinesHost = append(change, LinesHost[after:]...)
	WriteHostFile(LinesHost)
}

//FindEmptyLine
//AddIPBlock
//AddAlias

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
	fmt.Println("Hosts File ")
	fmt.Println("----------\n")
	for i := 0; i < len(LinesHost); i++ {
		fmt.Print(LinesHost[i])
	}
	WaitUser()
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
	WriteGroupNames()
	fmt.Print("Which group would you like to wiev: ")
	fmt.Scan(&nameGroup)
	control := 0
	temp := "# *" + nameGroup + "*"
	for i := 0; i < len(LinesHost); i++ {
		if strings.HasPrefix(LinesHost[i], "# *") && strings.Contains(LinesHost[i], temp) {
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

//Satır numarası ile komut satırı olmaktan çıkartıyor.
func RemoveCommendLineIP() {
	var lineNmbr int
	var newField string
	fmt.Printf("\n")
	for i := 0; i < len(LinesHost); i++ {
		fmt.Print((i + 1), " ", LinesHost[i])
	}
	fmt.Printf("\nLütfen satır numarası girin: ")
	fmt.Scan(&lineNmbr)
	newField = deleteCommendLine(lineNmbr - 1)
	AddLinesHosts(newField, (lineNmbr - 1), (lineNmbr))
}

//bu satır gerçekten ip satırı mı kontrolü ekle

func deleteCommendLine(lineNmbr int) string {
	var newField string
	temp := strings.Split(string(LinesHost[lineNmbr]), "#")
	//tempin boyutu kontrol edilmelidir
	newField = string(temp[1])
	return newField
}

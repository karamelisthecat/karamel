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
	fmt.Println("______________________________________")
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
		fmt.Println("Böyle bir grup bulunmamaktadır")
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
	var lineNmbr int    //kullanıcıdan alınan satır no
	var newField string // # çıkmış hali
	writeHostFilewithNmbr()
	lineNmbr = checkUserInput()
	newField = deleteCommendLine(lineNmbr - 1)
	AddLinesHosts(newField, (lineNmbr - 1), (lineNmbr))
	fmt.Println("______________________________________")
	fmt.Println("Success! Final version of hosts file: ")
	writeHostFilewithNmbr()
	WaitUser()
}

//kullanıcıdan satır numarasını alıyor ve o satır geçerli mi kontrol ediyor.
func checkUserInput() int {
	var lines string
	fmt.Printf("\nLütfen satır numarası girin: ")
	entry, _ := fmt.Scanf("%s", &lines)
	lineNmbr, err := strconv.Atoi(lines)
	if err != nil || entry == 0 || lineNmbr > len(LinesHost) || lineNmbr < 1 {
		fmt.Println("Geçersiz satır numarası girdiniz")
		checkUserInput()
	} else if strings.HasPrefix(LinesHost[lineNmbr-1], "#") != true || LinesHost[lineNmbr-1] == "" {
		fmt.Println("Bu satır yorum satırı değildir.")
		checkUserInput()
	} else if strings.HasPrefix(LinesHost[lineNmbr-1], "#") && strings.Contains(LinesHost[lineNmbr-1], "\t") != true {
		fmt.Println("Bu satır bir IP alanı içermez.")
		checkUserInput()
	}
	return lineNmbr
}

func deleteCommendLine(lineNmbr int) string {
	temp := strings.Split(string(LinesHost[lineNmbr]), "#")
	return string(temp[1])
}

func writeHostFilewithNmbr() {
	fmt.Printf("\n")
	for i := 0; i < len(LinesHost); i++ { //dosyayı yazdır
		fmt.Print((i + 1), " ", LinesHost[i])
	}
	fmt.Printf("\n")
}

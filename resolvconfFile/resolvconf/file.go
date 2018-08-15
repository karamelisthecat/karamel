package resolvconf

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	path                         = "/run/systemd/resolve/stub-resolv.conf"
	lines                        []string
	temp, change, tempNameserver []string
)

func OpenReadFile() {
	//resolv.conf dosyasını satır satır okuyor ve lines dizisine yazıyor.
	//yorum satırları dışındaki satırları ekrana basıyor.
	fmt.Println("----------RESOLV.CONF----------")
	file, _ := ioutil.ReadFile(path)
	lines = strings.SplitAfter(string(file), "\n")
	for i := 0; i < len(lines); i++ {
		control, j := 0, 0
		//yorum satırı kontrolü yapmak için satırdaki her bir karakteri oneLine'a ekliyor
		oneLine := strings.Split(lines[i], "")
		for ; j < len(oneLine); j++ {
			if oneLine[j] == "#" {
				control = 1
				break
			}
		}
		if control == 1 {
			if j != 0 {
				fmt.Printf((strings.Join(oneLine[:j], "")) + "\n")
			} else {
				continue
			}
		} else {

			fmt.Printf(strings.Join(oneLine, ""))
		}
	}
	fmt.Println("-------------------------------")

}
func KeepResolvconf() {
	//resolv.conf dosyasında nameserver alanıyla onun altında kalan alanları iki farklı arrayde turuyor
	var control, j int
	for j = 0; j < len(lines); j++ {
		if strings.Contains(lines[j], "nameserver ") == true {
			change = append(change, lines[:j]...)
			control = 1
			break
		}
	}
	if control == 1 {
		for ; j < len(lines); j++ {
			if strings.Contains(lines[j], "nameserver ") == false {
				temp = append(temp, lines[j])
			} else {
				tempNameserver = append(tempNameserver, lines[j])
			}
		}
	}

}
func SaveChange() {
	change = append(change, tempNameserver...)
	change = append(change, temp...)
	justString := strings.Join(change, "")
	bytes := []byte(justString)
	_ = ioutil.WriteFile(path, bytes, 0644)
	fmt.Println("LAST VIEW OF THE FILE")
	OpenReadFile()

}

package hostsfile

import (
	"fmt"
	"net"
	"strings"
)

// Adding ip block to /etc/hosts file.
func AddIPblock() {
	var addingField string
	var change []string
	addingField = ReturnField()
	for i := 0; i < len(LinesHost); i++ {
		if LinesHost[i] == "\n" {
			change = append(change, LinesHost[:i]...)    //iden öncesi
			change = append(change, addingField)         //alanı ekliyor
			LinesHost = append(change, LinesHost[i:]...) //birleştiriyor
			break
		}
	}
	WriteHostFile(LinesHost)
}

func ReturnField() string {
	var ipField string
	ipaddress, err := checkIP()
	if err != nil {
		fmt.Print("IP ADRESİ HATASI") //hata düzenlenecek.
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
		fmt.Print("\nEklemek istediğiniz hostname'i giriniz:") //sadece - ve . kontrolü eksik.
		fmt.Scan(&hostname)
		control, hostname = hostnameCheck(hostname)
	}
	return hostname
}

func hostnameCheck(hostname string) (bool, string) { // hostname daha önce girilmiş mi bakıyor
	for i := 0; i < len(LinesHost); i++ {
		if strings.Contains(LinesHost[i], hostname) && strings.Contains(LinesHost[i], "#") != true {
			fmt.Println(hostname, " :hostname daha önce kullanılmıştır. Yeni bir tane girin.")
			return true, hostname
		}
	}
	return false, hostname
}

func checkIP() (string, error) {
	var ipv4Addr net.IP
	for {
		var ipv4 string
		fmt.Printf("Eklemek istediğiniz IP adresini giriniz: ")
		fmt.Scanln(&ipv4)
		ipv4Addr = net.ParseIP(ipv4)
		if ipv4Addr.To4() == nil {
			fmt.Println("girdiğiniz ip adresi hatalıdır. ")
			//                  return nil, error.Error() err koy!
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
	var control = 0
	fmt.Print("Eklenecek grup adını girin: ")
	fmt.Scan(&group)
	for i := 0; i < len(GroupName); i++ {
		if GroupName[i] == group {
			fmt.Print("Bu grup bulunmaktadır") //err ekle
			control = 1
		}
	}
	if strings.Contains(group, "*") {
		control = 1
		fmt.Print("İçerisinde * işareti olan bir grup adı oluşturamazsınız")
	}
	if control == 0 {
		//      addNewLine() //grup eklemeden önce boşluk ekliyor
		newgroup = "# *" + group + "*" + "\n"
		LinesHost = append(LinesHost, newgroup)
		GroupName = append(GroupName, group)
		WriteHostFile(LinesHost)
		AddNewLine()
		setGroup(group)
	}
}

func AddFieldstoGroup() { //grubu alıyor kullanıcıdan
	var addGroup string
	fmt.Print("\nEklemek istediğiniz grup adını giriniz:")
	fmt.Scan(&addGroup)
	setGroup(addGroup)
}

func setGroup(addGroup string) { //grubu bulup araya alan ekliyor
	var control int
	var emptyLineTemp bool
	control = 0
	for i := 0; i < len(LinesHost); i++ { //satırlar arasında geziyor
		check, j := groupControl(i)
		if check && addGroup == GroupName[j] {
			emptyLineTemp = FindEmptyLine(i)
			if emptyLineTemp == true {
				control = 1
				break
			} //control = 0 da ne yapacak?
		}
		if control == 1 {
			WriteHostFile(LinesHost)
			break
		}
	}
}

func groupControl(i int) (bool, int) { //elemanı mı
	for j := 0; j < len(GroupName); j++ {
		if strings.HasPrefix(LinesHost[i], "# *") && strings.Contains(LinesHost[i], "*"+GroupName[j]+"*") {
			return true, j
		}
	}
	return false, -1
}

func AddAlias() {
	var change []string
	ipField, i := checkAlias()
	if i != -1 {
		change = append(change, LinesHost[:i]...)
		change = append(change, ipField)
		LinesHost = append(change, LinesHost[(i+1):]...)
		WriteHostFile(LinesHost)
	}
}

func checkAlias() (string, int) {
	var iptemp string
	var ctrl = false
	var ipField string
	var i int
	fmt.Print("\nHangi ip adresine alias eklensin: ")
	fmt.Scan(&iptemp)
	for i = 0; i < len(LinesHost); i++ {
		if strings.HasPrefix(LinesHost[i], iptemp) {
			fmt.Print("\nNe eklensin: ")
			var addTemp string
			fmt.Scan(&addTemp)
			bolTemp := strings.Split(string(LinesHost[i]), "\n")
			ipField = string(bolTemp[0]) + "\t" + addTemp + "\n"
			ctrl = true
			break
		}
	}
	if ctrl != true {
		fmt.Println("Böyle bir IP bulunmamaktadır")
		return iptemp, -1
	}
	return ipField, i
}

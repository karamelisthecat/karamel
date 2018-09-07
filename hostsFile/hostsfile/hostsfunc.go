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
	fmt.Print("Seçeneklerden birini seçiniz: ")
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
		fmt.Println("Geçersiz bir işlem girdiniz")
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
		fmt.Print("Eklemek istediğiniz hostname'i giriniz:") //sadece - ve . kontrolü eksik.
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
			fmt.Println("girdiğiniz ip adresi hatalıdır. \n")
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
	var control = true
	fmt.Print("Eklenecek grup adını girin: ")
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
		fmt.Printf("\n" + group + " isimli grup eklenmiştir!")
		FindtheGroup(group)
		LastViewoftheFile()
	}
}

//grup adı kullanılabilir mi diye bakıyor.
func groupCheck(group string) bool {
	for i := 0; i < len(GroupName); i++ {
		if GroupName[i] == group {
			fmt.Print("Bu grup bulunmaktadır") //err ekle
			return false
		}
	}
	if strings.Contains(group, "*") {
		fmt.Print("İçerisinde * işareti olan bir grup adı oluşturamazsınız")
		return false
	}
	return true
}

//Kullanıcıya daha fazla eklemek ister misiniz diye soruyor.
func askOpt() bool {
	var userOpt string
	fmt.Print("\nDaha fazla eklemek ister misiniz? \n(Eklemek istiyorsanız 'y' ya da 'Y' girmelisiniz): ")
	fmt.Scan(&userOpt)
	if userOpt == "y" || userOpt == "Y" {
		return true
	}
	return false
}

//Dosyanın son halini görüntülemek ister mi diye kullanıcıya soruyor.
func LastViewoftheFile() {
	var userOpt string
	fmt.Print("\nDosyanın son halini görüntülemek ister misiniz? \n(Görüntülemek için 'y' ya da 'Y' girmelisiniz: ")
	fmt.Scan(&userOpt)
	if userOpt == "y" || userOpt == "Y" {
		WriteLines()
	}
}

func AddFieldstoGroup() { //grubu alıyor kullanıcıdan
	var addGroup string
	//	fmt.Println("Bulunan Gruplar:")
	WriteGroupNames()
	fmt.Print("\nEklemek istediğiniz grup adını giriniz:")
	fmt.Scan(&addGroup)
	setGroup(addGroup)
	fmt.Println("Success!")
	LastViewoftheFile()
}

// Find the group and add ip field.
func setGroup(addGroup string) { //grubu bulup araya alan ekliyor
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
				AddLinesHosts(fieldTemp, c, c) //Burada dosyaya yazıyor.
				ctrl = 1
				break
			}
		}
	}
	if ctrl == 0 {
		fmt.Println("Böyle bir grup bulunmamaktadır.")
		WaitUser()
	}
}

//fonksiyon adını düzenle
func groupControl(i int) (bool, int) { //elemanı mı
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
	ipField, i := checkAlias()
	if i != -1 {
		AddLinesHosts(ipField, i, (i + 1))
		fmt.Println("Success!")
		LastViewoftheFile()
	} else {
		fmt.Printf("Bu ip'yi eklemek ister misiniz? \n(Eklemek istiyorsanız 'y' ya da 'Y' girmelisiniz): ")
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
	fmt.Printf("Hangi ip adresine alias eklensin: ")
	fmt.Scan(&iptemp)
	for i = 0; i < len(LinesHost); i++ {
		if strings.HasPrefix(LinesHost[i], iptemp) {
			fmt.Printf("Ne eklensin: ")
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

//İşleme devam etmeden önce kullanıcı girişini bekliyor.
func WaitUser() {
	var temp string
	entry, _ := fmt.Scanf("%s", &temp)
	if entry != 0 {
		fmt.Println("\n")
	}
}

// kullanıcı işlem sırasında iptal etmek isterse diye ??
func cancelOpt(ctrl string) {
	if ctrl == "q" {
		fmt.Println("Islem iptal edildi. \nAna menüye dönülüyor.")
		_ = UserOptMenu()
	}
}

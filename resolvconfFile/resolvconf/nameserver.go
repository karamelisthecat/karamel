package resolvconf

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	dns1 = &DNSName{"GoogleDNS", "8.8.8.8", "8.8.4.4"}
	dns2 = &DNSName{"OpenDNS", "208.67.222.222", "208.67.220.220"}
	dns3 = &DNSName{"Cloudflare", "1.1.1.1", "1.0.0.1"}
	dns4 = &DNSName{"Norton ConnectSafe", "199.85.126.10", "199.85.127.10"}
	dns5 = &DNSName{"Comodo Secure", "8.26.56.23", "8.20.247.20"}
	dns6 = &DNSName{"YandexDNS", "77.88.8.8", "77.88.8.1"}

	list  = [6]*DNSName{dns1, dns2, dns3, dns4, dns5, dns6}
	Addns *string
	Delns *int
)

type DNSName struct {
	name     string
	address1 string
	address2 string
}

func SelectMenu() {
	var choice string
	var dns string
	for {
		fmt.Println("1: Clear all and just add nameserver")
		fmt.Println("2: Modify file")
		fmt.Println("3: Exit")
		enter, _ := fmt.Scanf("%s", &choice)
		if enter == 0 {
			SelectMenu()
		}
		if choice == "1" {
			nameserverList()
			fmt.Printf("choice one or enter DNS: ")
			entry, _ := fmt.Scanf("%s", &dns)
			slct, _ := strconv.Atoi(dns)
			if entry == 0 {
				SelectMenu()
			}
			if 0 < slct && slct <= 6 {
				dns = selectNameserver(slct)
			}
			tempNameserver = tempNameserver[:0]
			tempNameserver = append(tempNameserver, "nameserver "+dns+"\n")

		} else if choice == "2" {
			nameserverMenu()
		} else if choice == "3" {
			break
		} else {
			SelectMenu()
		}
	}

}
func nameserverMenu() {
	//kullanıcıyı nameserver alanıyla ilgili yapmak istediği değişikliğe göre ilgili fonk yönlendiriyor.
	var choice string
	fmt.Println("1: Add nameserver")
	fmt.Println("2: Delete nameserver ")
	fmt.Println("3: Back to main menu ")
	enter, _ := fmt.Scanf("%s", &choice)
	if enter == 0 {
		nameserverMenu()
	}
	if choice == "1" {
		AddNameserver(Addns)
	} else if choice == "2" {
		DeleteNameserver(Delns)

	} else if choice == "3" {
		SelectMenu()
	} else {
		fmt.Println("please enter one of 2 options ")
		nameserverMenu()
	}
}

func nameserverList() {
	fmt.Println("**********************************************************")
	for i := 0; i < len(list); i++ {
		fmt.Println(i+1, ": ", list[i].name)
	}
	fmt.Println("**********************************************************")

}
func selectNameserver(slct int) (str string) {

	//çağrıldığı program için kullanıcının seçtiği dnsi return ediyor.
	return list[slct-1].address1
}

func AddNameserver(Addns *string) {
	var roww string
	var dns string
	for {
		if Addns != nil {
			//if strings.EqualFold(*Addns, "0.0.0.0") == false {
			var control = 0
			var dnsName int
			for dnsName = 0; dnsName < 6; dnsName++ {
				if strings.EqualFold(*Addns, list[dnsName].name) == true {
					tempNameserver[0] = "nameserver " + (list[dnsName].address1) + "\n"
					control = 1
					break
				}
			}
			if control == 0 {
				tempNameserver[0] = "nameserver " + *Addns + "\n"
			}
			break
		} else {

			fmt.Println("eklemek istediğiniz sırayı seçiniz")
			for i := 0; i < 3; i++ {
				if i < len(tempNameserver) {
					fmt.Println(i+1, ": ", tempNameserver[i])
				} else {
					fmt.Println(i+1, ":\n")
				}
			}
			ent, _ := fmt.Scanf("%s", &roww)
			row, _ := strconv.Atoi(roww)
			if ent == 0 || row/1 != row || row > 3 {
				break
			}
			fmt.Println("choose one or enter different DNS: ")
			nameserverList()
			entry, _ := fmt.Scanf("%s", &dns)
			slct, _ := strconv.Atoi(dns)
			if entry == 0 {
				break
			}

			if 0 < slct && slct <= 6 {
				dns = selectNameserver(slct)
			}
			if row < len(tempNameserver) {
				tempNameserver[row-1] = "nameserver " + dns + "\n"
			} else {
				tempNameserver = append(tempNameserver, "nameserver "+dns+"\n")
			}
		}
	}
}
func DeleteNameserver(Delns *int) {
	for {
		if Delns != nil {
			if *Delns >= 1 && *Delns <= 3 {

				tempNameserver[len(tempNameserver)-1] = tempNameserver[*Delns-1]
				tempNameserver[*Delns-1] = tempNameserver[len(tempNameserver)-1]
				tempNameserver = tempNameserver[:len(tempNameserver)-1]

			}
			break
		} else {

			var no string

			for i := 0; i < len(tempNameserver); i++ {
				fmt.Println(i+1, ":", tempNameserver[i])
			}
			fmt.Printf("please select the number you want to delete in nameservers: ")

			entry, _ := fmt.Scanf("%s", &no)
			if entry == 0 {
				SelectMenu()
			}
			number, _ := strconv.Atoi(no)
			number = number - 1
			tempNameserver[len(tempNameserver)-1] = tempNameserver[number]
			tempNameserver[number] = tempNameserver[len(tempNameserver)-1]
			tempNameserver = tempNameserver[:len(tempNameserver)-1]

		}
	}

}

package resolvconf

import (
	"fmt"
	"net"
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
Loop:
	for {
		fmt.Println("1: Clear all and just add nameserver")
		fmt.Println("2: Modify file")
		fmt.Println("3: List nameservers")
		fmt.Println("q: Exit")
		enter, _ := fmt.Scanf("%s", &choice)
		if enter == 0 {
			SelectMenu()
		}
		switch choice {
		case "1":
			addOneNameserver()
			break
		case "2":
			nameserverMenu()
			break
		case "3":
			lastViewNameserver()
			break
		case "q":
			break Loop
		default:
			fmt.Println("Please enter one of 3 options!")
			continue Loop
		}
	}
}
func nameserverMenu() {
	//kullanıcıyı nameserver alanıyla ilgili yapmak istediği değişikliğe göre ilgili fonk yönlendiriyor.
Loop:
	for {
		var choice string
		fmt.Println("1: Add nameserver")
		fmt.Println("2: Delete nameserver ")
		fmt.Println("q: Back to main menu ")
		enter, _ := fmt.Scanf("%s", &choice)
		if enter == 0 {
			nameserverMenu()
		}
		switch choice {
		case "1":
			AddNameserver(Addns)
			break Loop
		case "2":
			DeleteNameserver(Delns)
			break Loop
		case "q":
			break Loop
		default:
			fmt.Println("please enter one of 3 options! ")
			continue Loop
		}
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
func addOneNameserver() {
	var dns string
	nameserverList()
	fmt.Printf("choice one or enter DNS(enter 'q' for exit): ")
	entry, _ := fmt.Scanf("%s", &dns)
	if dns == "q" {
		SelectMenu()
	}
	slct, _ := strconv.Atoi(dns)
	if entry == 0 {
		SelectMenu()
	}
	if 0 < slct && slct <= 6 {
		dns = selectNameserver(slct)
	}
	if controlNameserver(dns) == true {
		tempNameserver = tempNameserver[:0]
		tempNameserver = append(tempNameserver, "nameserver "+dns+"\n")
		fmt.Println("\nLAST VIEW OF THE NAMESERVER LIST")
		fmt.Println("**********************************************************")
		for i := 0; i < len(tempNameserver); i++ {
			fmt.Println(tempNameserver[i])
		}
		fmt.Println("**********************************************************")
	} else {
		fmt.Println("invalid IP address")
		addOneNameserver()
	}
}
func Adding(dns string) {
	tempNameserver = tempNameserver[:0]
	tempNameserver = append(tempNameserver, "nameserver "+dns+"\n")
}

func lastViewNameserver() {
	fmt.Println("**********************************************************")
	for i := 0; i < len(tempNameserver); i++ {
		fmt.Println(tempNameserver[i])
	}
	fmt.Println("**********************************************************")

}

func AddNameserver(Addns *string) {
	var rowStr string
	var dns string
	for {
		if Addns != nil {
			addNameserverFlag(Addns)
			break
		} else {
			fmt.Println("Select the line or enter 'q' for exit")
			for i := 0; i < 3; i++ {
				if i < len(tempNameserver) {
					fmt.Println(i+1, ": ", tempNameserver[i])
				} else {
					fmt.Println(i+1, ":\n")
				}
			}
			ent, _ := fmt.Scanf("%s", &rowStr)
			if rowStr == "q" || ent == 0 {
				break
			}
			row, _ := strconv.Atoi(rowStr)
			if row != 1 && row != 2 && row != 3 {
				continue
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
			if controlNameserver(dns) == true {
				AddingRow(row, dns)
			} else {
				fmt.Println("invalid IP address")
				continue
			}
		}
	}
}
func AddingRow(row int, dns string) {
	if row <= len(tempNameserver) {
		tempNameserver[row-1] = "nameserver " + dns + "\n"
	} else {
		tempNameserver = append(tempNameserver, "nameserver "+dns+"\n")
	}

}

func addNameserverFlag(Addns *string) {
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
			listNameserver()
			fmt.Printf("please select the number you want to delete in nameservers or enter 'q' for exit: ")
			entry, _ := fmt.Scanf("%s", &no)
			if no == "q" {
				break
			}
			if entry == 0 {
				SelectMenu()
			}
			number, _ := strconv.Atoi(no)
			if number > len(tempNameserver) {
				continue
			}
			if number != 1 && number != 2 && number != 3 {
				continue
			}
			Deleting(number)
		}
	}
}
func listNameserver() {

	for i := 0; i < len(tempNameserver); i++ {
		fmt.Println(i+1, ":", tempNameserver[i])
	}

}
func Deleting(number int) {
	number = number

	tempNameserver[number-1] = tempNameserver[len(tempNameserver)-1]

	tempNameserver[len(tempNameserver)-1] = tempNameserver[number-1]
	//tempNameserver[number-1] = tempNameserver[len(tempNameserver)-1]
	tempNameserver = tempNameserver[:len(tempNameserver)-1]

}
func controlNameserver(nameserver string) bool {
	var ipv4Addr net.IP
	var control = true
	for {
		ips, err := net.LookupIP(nameserver)
		if err != nil {
			//fmt.Println("Invalid IP address. ")
			control = false
			break
		}
		yeni := ips[0].String()
		ipv4Addr = net.ParseIP(yeni)
		if ipv4Addr == nil {
			//	fmt.Println("the IP address you entered is incorrect.")
			control = false
			break
		}
		kontrol := ipv4Addr.DefaultMask()
		if kontrol == nil {
			//	fmt.Println("Invalid IP address. ")
			control = false
			break
		}
		break
	}
	return control
}

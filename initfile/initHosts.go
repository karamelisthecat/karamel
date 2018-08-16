package initfile

import (
	"fmt"
	"github.com/karamelisthecat/karamel/hostsFile/flag"
	"github.com/karamelisthecat/karamel/hostsFile/hostsfile"
)

func InitHosts() {
	hostsfile.LinesHost, _ = hostsfile.ReadHostFile("/etc/hosts")
	hostsfile.FindGroupNames()
	flag.OneFlag()
	userInterface()
}

func userInterface() {
	var isRunning bool
	var optionNumber string
	isRunning = true
	fmt.Println("\n/etc/hosts file")
	fmt.Println("----------------")
	for isRunning {
		fmt.Println("______________________________________")
		fmt.Println("1. grup ekleyiniz.\n2. grup görüntüleyiniz.\n3. bulunan grupların listesi.\n4. ip adresi ekleyiniz.\n5. dosyayı yazdır\n6. gruba alan ekle.\n7. IP adresine alias ekle.\nAna menüye dönmek için 'q'ya basınız.")
		fmt.Println("______________________________________")
		fmt.Print("\nSeçeneklerden birini seçiniz: ")
		fmt.Scan(&optionNumber)
		fmt.Printf("\n")
		switch optionNumber {
		case "q":
			fmt.Println("Ana menüye dönülüyor.")
			isRunning = false
			break
		case "1":
			hostsfile.AddGroup()
		case "2":
			hostsfile.ListGroup()
		case "3":
			hostsfile.WriteGroupNames()
		case "4":
			hostsfile.AddIPblock()
		case "5":
			hostsfile.WriteLines()
		case "6":
			hostsfile.AddFieldstoGroup()
		case "7":
			hostsfile.AddAlias()
		default:
			fmt.Println("Geçersiz bir işlem girdiniz")
		}
	}
}
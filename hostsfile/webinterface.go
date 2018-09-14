package hostsfile

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func WebInterface() {
	http.HandleFunc("/addip", addip)
	http.HandleFunc("/addgroup", addGroup)
	http.HandleFunc("/listhostsfile", listhostsfile)
	http.HandleFunc("/addalias", addalias)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func addalias(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addalias method:", r.Method) //get request method
	if r.Method == "GET" {
		t, err := template.ParseFiles("./webinterface/addalias.gtpl")
		fmt.Println(err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("IPAddress:", r.Form["ipaddress"])
		fmt.Println("Alias", r.Form["alias"])

		ipaddressTemp := strings.Join(r.Form["ipaddress"], "")
		aliasTemp := strings.Join(r.Form["alias"], "")
		AddAliasInterface(ipaddressTemp, aliasTemp)
	}
}

func addip(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addip method:", r.Method) //get request method
	if r.Method == "GET" {
		t, err := template.ParseFiles("./webinterface/addip.gtpl")
		fmt.Println(err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("IPAddress:", r.Form["ipaddress"])
		fmt.Println("Hostname:", r.Form["hostname"])

		ipaddressTemp := strings.Join(r.Form["ipaddress"], "")
		hostnameTemp := strings.Join(r.Form["hostname"], "")
		ReturnipField(ipaddressTemp, hostnameTemp)
	}
}

func addGroup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./webinterface/addgroup.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("GroupName:", r.Form["groupname"])
		fmt.Println("IPAddress:", r.Form["ipaddress"])
		fmt.Println("Hostname:", r.Form["hostname"])

		groupnameTemp := strings.Join(r.Form["groupname"], "")
		ipaddressTemp := strings.Join(r.Form["ipaddress"], "")
		hostnameTemp := strings.Join(r.Form["hostname"], "")
		AddGroupInterface(groupnameTemp, ipaddressTemp, hostnameTemp)
	}
}

func listhostsfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hosts File: ")
	for i := 0; i < len(LinesHost); i++ {
		fmt.Fprintf(w, LinesHost[i])

	}
}

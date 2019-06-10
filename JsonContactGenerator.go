package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ContactsList struct {
	Version  string        `json:"version"`
	Contacts []ContactData `json:"contacts"`
}

type ContactData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var path = "./mock-contacts.json"

func createFile(data string) {
	var file, err = os.Create(path)
	fmt.Println(err)
	file.WriteString(data)
	defer file.Close()
	fmt.Println("File writed Successfully", path)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter max of contacts: ")
	inputMax, _ := reader.ReadString('\n')
	pureInputMax := (strings.TrimSuffix(inputMax, "\n"))
	MAX, _ := strconv.Atoi(pureInputMax)

	var contactsList ContactsList
	contactsList.Version = "1"

	for i := 1; i <= MAX; i++ {
		var num = strconv.Itoa(i)
		if len(strconv.Itoa(MAX)) > len(strconv.Itoa(i)) {
			addZero := len(strconv.Itoa(MAX)) - len(strconv.Itoa(i))
			for p := 0; p < addZero; p++ {
				num = "0" + num
			}
		}
		contactsList.Contacts = append(contactsList.Contacts, ContactData{Name: "Name" + string(num), Email: "Name" + string(num) + "@testmail.com"})
	}

	output, _ := json.MarshalIndent(&contactsList, "", "  ")
	fmt.Println(string(output))
	createFile(string(output))
}

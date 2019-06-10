package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "./mock-contacts.vcf"

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

	var buffer bytes.Buffer

	for i := 1; i <= MAX; i++ {
		var num = strconv.Itoa(i)
		if len(strconv.Itoa(MAX)) > len(strconv.Itoa(i)) {
			addZero := len(strconv.Itoa(MAX)) - len(strconv.Itoa(i))
			for p := 0; p < addZero; p++ {
				num = "0" + num
			}
		}
		buffer.WriteString("BEGIN:VCARD\n")
		buffer.WriteString("VERSION:3.0\n")
		buffer.WriteString("FN:Test" + num + "\n")
		buffer.WriteString("EMAIL;TYPE=HOME,INTERNET,pref:Test" + num + "@gmail.com\n")
		buffer.WriteString("END:VCARD\n")
	}

	fmt.Printf(buffer.String())
	createFile(buffer.String())
}

package _io

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func writeJsonFile() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	js, err := json.Marshal(vc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("JSON format: %s", js)
	file, err := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(vc)
	if err != nil {
		fmt.Println(err)
	}

}

func readJsonFile() {
	file, err := os.Open("vcard.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	reader := bufio.NewReader(file)
	sb := strings.Builder{}

	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		sb.WriteString(readString)
	}

	vcard := VCard{}
	err = json.Unmarshal([]byte(sb.String()), &vcard)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("vcard=== %+v", vcard)
}

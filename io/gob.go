package _io

import (
	"encoding/gob"
	"fmt"
	"os"
)

type P struct {
	X, Y, Z int
	Name    string
}

func writeGob() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	file, err := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(vc)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func readGob() {
	file, err := os.Open("vcard.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	decoder := gob.NewDecoder(file)
	var vcard = VCard{}
	err = decoder.Decode(&vcard)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vcard)
}

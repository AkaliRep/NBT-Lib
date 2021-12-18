package main

import (
	"bytes"
	"fmt"
	"log"
	"nbtlib/nbt"
)

func main() {
	file, err := nbt.DecodeFile("bigtest.nbt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", file)

	r := bytes.NewReader(file)

	nbtCompound := nbt.Read(r)
	log.Printf("%v\n", nbtCompound)
}

package main

import (
	"log"
	"nbtlib/nbt"
)

func main() {
	nbtCompound := nbt.ReadFile("bigtest.nbt")
	log.Printf("%v\n", nbtCompound)
}

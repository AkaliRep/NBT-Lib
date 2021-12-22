package main

import (
	"fmt"
	"nbtlib/nbt"
)

func main() {
	fmt.Printf("%s", nbt.ReadFile("bigtest.nbt"))
}

package main

import (
	"fmt"
	"nbtlib/nbt"
)

func main() {
	base := nbt.ReadFile("level.dat").GetCompound("Data")

	fmt.Printf("%s\n", base)
}

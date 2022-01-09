package main

import (
	"fmt"
	"log"
	"nbtlib/nbt"
	"os"
)

func main() {
	base := nbt.ReadFile("level.dat")
	result := make([]byte, 0)
	nbt.WriteImplicitCompound(&result, base)

	fmt.Printf("Converted: %v\n", result)

	f, err := nbt.DecompressFile("level.dat")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Original: %v\n", f)

	fmt.Printf("%s", nbt.ReadByteArray(result))
	fmt.Printf("%s", base)

	file, err := os.Create("output.nbt")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(result)
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}

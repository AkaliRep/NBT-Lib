package main

import (
	"fmt"
	"nbtlib/nbt"
)

type Chunk struct {
	X    int32  `nbt_name:"X"`
	Name string `nbt_name:"still"`
}

type World struct {
	C    Chunk  `nbt_name:"chunk"`
	Name string `nbt_name:"str"`
}

func main() {
	tag := nbt.ReadFile("hello_world.nbt")
	c := World{}
	err := nbt.FillStruct(&c, tag)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", c)
}

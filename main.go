package main

import (
	"fmt"
	testmodML "github.com/yeqf911/testmod/v2"
	"github.com/yeqf911/testmod"
)

func main() {
	g, err := testmodML.Hi("Alan", "fr")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
	fmt.Println(testmod.Hi("Bob"))
}

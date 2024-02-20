package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Rene"
	var namePointer *string
	namePointer = &name

	fmt.Println("NamePointer address:", namePointer)
	fmt.Println("NamePointer value:", *namePointer)

	fmt.Println("Simple Name:", name)
	fmt.Println("Decorate Name: ", addDecoration(name))
	fmt.Println("Decorate Name1: ", addDecoration1(&name))
	addDecoration2(namePointer)
	fmt.Println("Decorate Name2: ", name)
}

func addDecoration(name string) string {
	newName := "( ( ( " + strings.ToUpper(name) + " ) ) )"
	return newName
}

func addDecoration1(name *string) string {
	newName := "1( ( ( " + strings.ToUpper(*name) + " ) ) )"
	return newName
}

func addDecoration2(name *string) {
	*name = "2( ( ( " + strings.ToUpper(*name) + " ) ) )"
}

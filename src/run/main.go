package main

import (
	"../wireup"
	"fmt"
	"log"
)

func main() {
	em := ecsinitialize.InitializeECS()
	log.Println(em)
	var input string
	fmt.Scanln(&input)
}

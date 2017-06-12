package main

import (
	"log"
	"wireup"
)

func main() {
	em := ecsinitialize.InitializeECS()
	log.Println(em)
	var input string
	log.Scanln(&input)
}

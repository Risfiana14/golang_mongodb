package main

import (
	"fmt"
	"log"
	"tugas8/utils"
)

func main() {
	password := "123456" // password user baru
	hash, err := utils.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash password baru:", hash)
}
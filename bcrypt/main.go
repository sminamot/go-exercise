package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := []byte("password")

	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(hash))

	err = bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		log.Fatal(err)
	}
}

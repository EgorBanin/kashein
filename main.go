package main

import (
	"kashein/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

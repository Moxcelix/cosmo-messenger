package main

import (
	"log"
	"main/cmd"
)

func main() {
	if err := cmd.StartApp(); err != nil {
		log.Fatal(err)
	}
}

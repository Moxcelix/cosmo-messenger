package main

import (
	"log"
	"messenger/cmd"
)

func main() {
	if err := cmd.StartApp(); err != nil {
		log.Fatal(err)
	}
}

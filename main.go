package main

// @title Cosmo Messenger API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @security BearerAuth

import (
	"log"
	"main/cmd"
	_ "main/docs"
)

func main() {
	if err := cmd.StartApp(); err != nil {
		log.Fatal(err)
	}
}

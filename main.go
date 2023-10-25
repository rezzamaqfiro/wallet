package main

import (
	"log"

	"github.com/rezzamaqfiro/wallet/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

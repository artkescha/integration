package main

import (
	"log"

	"github.com/artkescha/integration/cmd"
)

func main() {
	if err := cmd.Command.Execute(); err != nil {
		log.Fatal(err)
	}
}

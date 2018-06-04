package main

import (
	"log"
	"os"

	"github.com/matthewmueller/pogo/cli"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

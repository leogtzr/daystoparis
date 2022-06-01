package main

import (
	"os"

	"github.com/leogtzr/daystoparis/daystoparis"
)

func main() {
	os.Exit(daystoparis.CLI(os.Args[:]))
}

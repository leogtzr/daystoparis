package main

import (
	"os"

	_ "github.com/leogtzr/daystoparis/daystoparis"
)

func main() {
	os.Exit(daystoparis.CLI(os.Args[:]))
}

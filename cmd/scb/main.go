package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("(help message goes here)\n")
}

func main() {
	var sc string
	var err error
	sc, os.Args, err = extractSubcommand(os.Args)

	if err != nil {
		usage()
		return
	}

	err = runSubcommand(sc)
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		usage()
		return
	}
}

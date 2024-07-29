package main

import (
	"flag"
	"fmt"
)

var args []string

func usage() {
	fmt.Printf("(help message goes here)\n")
}

var into *string

func main() {
	into = flag.String("into", ".", "put files in this dir")
	flag.Parse()

	var sc string
	var err error
	sc, args, err = extractSubcommand(flag.Args())

	if err != nil {
		usage()
		return
	}

	err = runSubcommand(sc, args)
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		usage()
		return
	}
}
